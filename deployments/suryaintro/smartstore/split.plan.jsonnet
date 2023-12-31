// Code generated by stencil-smartstore. DO NOT EDIT.
// This file generates a pod for split shard plan generation
local ok = import 'kubernetes/outreach.libsonnet';
local app = (import 'kubernetes/app.libsonnet').info('suryaintro');
local appImageRegistry = std.extVar('appImageRegistry');
local name = '%s-split-apply' % app.name;

local pgvolumes = import 'postgresql.volumes.jsonnet';
local resourceName = std.extVar('resourceName');
// the pod created below expects to be in the same namespace as
// smartstore deployment, it shares the service account, roles,
// configmaps, and secrets with a smartstore based deployment.
local configVolumeMounts = {
  ['config-pg-%s-split-plan-volume' % resourceName]: {
    mountPath: '/run/config/outreach.io/plan.yaml',
    subPath: 'plan.yaml',
  },
};
local configVolume = {
  ['config-pg-%s-split-plan-volume' % resourceName]: ok.ConfigMapVolume(ok.ConfigMap('pg-%s-split-plan-cfg' % resourceName, app.namespace)),
};
local input = std.split(std.extVar('input'), ';');
local getSourceDestMap(x) =
  {
    src_shard_key: std.split(x, ':')[0],
    dest_shard_key: std.split(x, ':')[1],
  };
local pod = {
  ['pg-%s-split-plan-cfg' % resourceName]: ok.ConfigMap('pg-%s-split-plan-cfg' % resourceName, app.namespace) {
    local this = self,
    data_:: {
      name: 'split-plan-%s' % std.extVar('planName'),
      schema_prefix: std.extVar('schemaPrefix'),
      plan_type: std.extVar('planType'),
      plan_spec:
        if std.extVar('planType') != 'move-shard' then {} else {
          entries: [
            getSourceDestMap(x)
            for x in input
          ],
        },
    },
    data: {
      'plan.yaml': std.manifestYamlDoc(this.data_),
    },
  },
  pod: ok.Pod('%s-split-plan' % app.name) {
    local this = self,
    metadata+: {
      [if app.namespace != null then 'namespace']: app.namespace,
      labels+: {
        repo: app.name,
        bento: app.bento,
        reporting_team: 'cad-fri',
      },
      annotations+: {
        'iam.amazonaws.com/role': '%s_service_role' % app.name,
        // https://docs.datadoghq.com/integrations/openmetrics/
        ['ad.datadoghq.com/' + name + '.check_names']: '["openmetrics"]',
        ['ad.datadoghq.com/' + name + '.init_configs']: '[{}]',
        ['ad.datadoghq.com/' + name + '.instances']: std.manifestJsonEx(self.datadog_prom_instances_, '  '),
        datadog_prom_instances_:: [
          {
            prometheus_url: 'http://%%host%%:8000' +
                            '/metrics',
            namespace: name,
            metrics: ['*'],
            send_distribution_buckets: true,
          },
        ],
      },
    },
    spec+: {
      // shares service account with owning service/deployment
      serviceAccountName: '%s-svc' % app.name,
      priorityClassName: 'high-priority',
      restartPolicy: 'Never',
      containers_:: {
        default: ok.Container(name) {
          image: '%s/%s:%s' % [appImageRegistry, app.name, app.version],
          imagePullPolicy: 'IfNotPresent',
          command: ['/usr/local/bin/smartstore'],
          args: [
            '--skip-update',
            'split',
            'plan',
            '--file=/run/config/outreach.io/plan.yaml',
            '--managed-resource-name=%s' % resourceName,
            '--schema-prefix=%s' % std.extVar('schemaPrefix'),
          ],
          env_+:: {
            MY_POD_SERVICE_ACCOUNT: ok.FieldRef('spec.serviceAccountName'),
            MY_NAMESPACE: ok.FieldRef('metadata.namespace'),
            MY_POD_NAME: name,
            MY_NODE_NAME: ok.FieldRef('spec.nodeName'),
            MY_DEPLOYMENT: app.name,
            MY_ENVIRONMENT: app.environment,
            MY_CLUSTER: app.cluster,
          },
          resources: {
            requests: {
              cpu: '1',
              memory: '512Mi',
            },
            limits: {
              cpu: '1',
              memory: '1000Mi',
            },
          },
          volumeMounts_+:: pgvolumes.postgresql_volume_mounts + configVolumeMounts,
        },
      },
      volumes_+:: pgvolumes.default_volumes + pgvolumes.postgresql_volumes + configVolume,
    },
  },
};

pod
