// Code generated by stencil-smartstore. DO NOT EDIT.
local ok = import 'kubernetes/outreach.libsonnet';
local app = (import 'kubernetes/app.libsonnet').info('suryaintro');
local appImageRegistry = std.extVar('appImageRegistry');
local pgvolumes = import 'postgresql.volumes.jsonnet';
local resources = import '../resources.libsonnet';

{
  SharedLabels: {
    repo: app.name,
    bento: app.bento,
    reporting_team: 'cad-fri',
  },
  // Renders a Kubernetes deployment with defaults common to all smartstore Deployments
  Deployment(name, namespace): ok.Deployment(name, namespace) {
    local ports = {
      'http-prom': { containerPort: 8000 },
      grpc: { containerPort: 5000 },
    },

    metadata+: {
      labels+: $.SharedLabels,
    },
    spec+: {
      replicas: 1,
      template+: {
        metadata+: {
          labels+: $.SharedLabels,
          annotations+: {
            'iam.amazonaws.com/role': '%s_service_role' % name,
            // https://docs.datadoghq.com/integrations/openmetrics/
            ['ad.datadoghq.com/' + name + '.check_names']: '["openmetrics"]',
            ['ad.datadoghq.com/' + name + '.init_configs']: '[{}]',
            ['ad.datadoghq.com/' + name + '.instances']: std.manifestJsonEx(self.datadog_prom_instances_, '  '),
            datadog_prom_instances_:: [
              {
                prometheus_url: 'http://%%host%%:' +
                                ports['http-prom'].containerPort +
                                '/metrics',
                namespace: name,
                metrics: ['*'],
                send_distribution_buckets: true,
              },
            ],
          },
        },
        spec+: {
          priorityClassName: 'high-priority',
          containers_:: {
            default: ok.Container(name) {
              image: '%s/%s:%s' % [appImageRegistry, app.name, app.version],
              // We don't want to ever pull the same tag multiple times.
              // In dev, this is replaced by sharing docker image cache with Kubernetes
              // so we also don't need to pull images.
              imagePullPolicy: 'IfNotPresent',
              env_+:: {
                MY_POD_SERVICE_ACCOUNT: ok.FieldRef('spec.serviceAccountName'),
                MY_NAMESPACE: ok.FieldRef('metadata.namespace'),
                MY_POD_NAME: ok.FieldRef('metadata.name'),
                MY_NODE_NAME: ok.FieldRef('spec.nodeName'),
                MY_DEPLOYMENT: name,
                MY_ENVIRONMENT: app.environment,
                MY_CLUSTER: app.cluster,
              },
              ports_+:: ports,
              resources: resources,
              volumeMounts_+:: pgvolumes.default_volume_mounts,
            },
          },
          volumes_+:: pgvolumes.default_volumes,
        },
      },
    },
  },
}
