// Code generated by stencil-smartstore. DO NOT EDIT.
// This file generates a pod for split shard plan generation
local ok = import 'kubernetes/outreach.libsonnet';
local app = (import 'kubernetes/app.libsonnet').info('suryaintro');
local appImageRegistry = std.extVar('appImageRegistry');

local pgvolumes = import 'postgresql.volumes.jsonnet';
local resourceName = std.extVar('resourceName');
local name = '%s-split-apply' % app.name;
// the pod created below expects to be in the same namespace as
// smartstore deployment, it shares the service account, roles,
// configmaps, and secrets with a smartstore based deployment.
local pod = {
  pod: ok.Pod(name) {
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
            'apply',
            '--managed-resource-name=%s' % resourceName,
            '--plan-name=%s' % std.extVar('planName'),
          ],
          env_+:: {
            // the postgresql client utilities that apply uses requires explicitly
            // locating the root.crt for rds CA. This env var, PGSSLROOTCERT, ensures
            // processes like psql and pg_dump are able to locate the CA certs.
            PGSSLROOTCERT: '/usr/local/share/ca-certificates/rds-combined-ca-bundle.pem',
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
          volumeMounts_+:: pgvolumes.postgresql_volume_mounts,
        },
      },
      volumes_+:: pgvolumes.default_volumes + pgvolumes.postgresql_volumes,
    },
  },
};

pod
