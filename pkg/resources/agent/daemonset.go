package agent

import (
	"crypto/sha256"
	"fmt"
	"strconv"

	"github.com/banzaicloud/operator-tools/pkg/reconciler"
	util "github.com/banzaicloud/operator-tools/pkg/utils"
	"github.com/spaghettifunk/vector-operator/pkg/resources/templates"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

const (
	BufferStorageVolume = "buffers"
)

func (r *Reconciler) daemonSet() (runtime.Object, reconciler.DesiredState, error) {

	labels := util.MergeLabels(r.Vector.Spec.AgentSpec.Labels, r.getAgentLabels())
	meta := r.AgentObjectMeta(agentDaemonSetName)
	meta.Annotations = util.MergeLabels(meta.Annotations, r.Vector.Spec.AgentSpec.DaemonSetAnnotations)
	podMeta := metav1.ObjectMeta{
		Labels:      labels,
		Annotations: r.Vector.Spec.AgentSpec.Annotations,
	}

	if r.configs != nil {
		for key, config := range r.configs {
			h := sha256.New()
			_, _ = h.Write(config)
			podMeta = templates.Annotate(podMeta, fmt.Sprintf("checksum/%s", key), fmt.Sprintf("%x", h.Sum(nil)))
		}
	}

	containers := []corev1.Container{
		*r.agentContainer(),
	}

	desired := &appsv1.DaemonSet{
		ObjectMeta: meta,
		Spec: appsv1.DaemonSetSpec{
			Selector: &metav1.LabelSelector{MatchLabels: util.MergeLabels(r.Vector.Spec.AgentSpec.Labels, r.getAgentLabels())},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: podMeta,
				Spec: corev1.PodSpec{
					ServiceAccountName: r.getServiceAccount(),
					Volumes:            r.generateVolume(),
					Tolerations:        r.Vector.Spec.AgentSpec.Tolerations,
					NodeSelector:       r.Vector.Spec.AgentSpec.NodeSelector,
					Affinity:           r.Vector.Spec.AgentSpec.Affinity,
					ImagePullSecrets:   r.Vector.Spec.AgentSpec.Image.ImagePullSecrets,

					Containers: containers,
				},
			},
		},
	}

	fmt.Println(desired)

	return desired, reconciler.StatePresent, nil
}

func (r *Reconciler) agentContainer() *corev1.Container {
	return &corev1.Container{
		Name:            containerName,
		Image:           r.Vector.Spec.AgentSpec.Image.RepositoryWithTag(),
		ImagePullPolicy: corev1.PullPolicy(r.Vector.Spec.AgentSpec.Image.PullPolicy),
		Ports:           r.generatePortsMetrics(),
		Resources:       r.Vector.Spec.AgentSpec.Resources,
		VolumeMounts:    r.generateVolumeMounts(),
		SecurityContext: &corev1.SecurityContext{
			RunAsUser:                r.Vector.Spec.AgentSpec.Security.SecurityContext.RunAsUser,
			RunAsNonRoot:             r.Vector.Spec.AgentSpec.Security.SecurityContext.RunAsNonRoot,
			ReadOnlyRootFilesystem:   r.Vector.Spec.AgentSpec.Security.SecurityContext.ReadOnlyRootFilesystem,
			AllowPrivilegeEscalation: r.Vector.Spec.AgentSpec.Security.SecurityContext.AllowPrivilegeEscalation,
			Privileged:               r.Vector.Spec.AgentSpec.Security.SecurityContext.Privileged,
			SELinuxOptions:           r.Vector.Spec.AgentSpec.Security.SecurityContext.SELinuxOptions,
		},
		Env:            r.Vector.Spec.AgentSpec.EnvVars,
		LivenessProbe:  r.Vector.Spec.AgentSpec.LivenessProbe,
		ReadinessProbe: r.Vector.Spec.AgentSpec.ReadinessProbe,
	}
}

func (r *Reconciler) generatePortsMetrics() (containerPorts []corev1.ContainerPort) {
	if r.Vector.Spec.AgentSpec.Metrics != nil && r.Vector.Spec.AgentSpec.Metrics.Port != 0 {
		containerPorts = append(containerPorts, corev1.ContainerPort{
			Name:          "monitor",
			ContainerPort: r.Vector.Spec.AgentSpec.Metrics.Port,
			Protocol:      corev1.ProtocolTCP,
		})
	}
	return
}

func (r *Reconciler) generateVolumeMounts() (v []corev1.VolumeMount) {
	v = []corev1.VolumeMount{
		{
			Name:      "varlibcontainers",
			ReadOnly:  true,
			MountPath: "/var/lib/docker/containers",
		},
		{
			Name:      "varlogs",
			ReadOnly:  true,
			MountPath: "/var/log/",
		},
	}

	for vCount, vMnt := range r.Vector.Spec.AgentSpec.ExtraVolumeMounts {
		v = append(v, corev1.VolumeMount{
			Name:      "extravolumemount" + strconv.Itoa(vCount),
			ReadOnly:  *vMnt.ReadOnly,
			MountPath: vMnt.Destination,
		})
	}

	// if r.Vector.Spec.AgentSpec.CustomConfigSecret == "" {
	// 	v = append(v, corev1.VolumeMount{
	// 		Name:      "config",
	// 		MountPath: "/fluent-bit/etc/fluent-bit.conf",
	// 		SubPath:   BaseConfigName,
	// 	})
	// } else {
	// 	v = append(v, corev1.VolumeMount{
	// 		Name:      "config",
	// 		MountPath: "/fluent-bit/etc/",
	// 	})
	// }

	return
}

func (r *Reconciler) generateVolume() (v []corev1.Volume) {
	v = []corev1.Volume{
		{
			Name: "varlibcontainers",
			VolumeSource: corev1.VolumeSource{
				HostPath: &corev1.HostPathVolumeSource{
					Path: r.Vector.Spec.AgentSpec.MountPath,
				},
			},
		},
		{
			Name: "varlogs",
			VolumeSource: corev1.VolumeSource{
				HostPath: &corev1.HostPathVolumeSource{
					Path: "/var/log",
				},
			},
		},
	}

	for vCount, vMnt := range r.Vector.Spec.AgentSpec.ExtraVolumeMounts {
		v = append(v, corev1.Volume{
			Name: "extravolumemount" + strconv.Itoa(vCount),
			VolumeSource: corev1.VolumeSource{
				HostPath: &corev1.HostPathVolumeSource{
					Path: vMnt.Source,
				},
			}})
	}

	// if r.Vector.Spec.AgentSpec.CustomConfigSecret == "" {
	// 	volume := corev1.Volume{
	// 		Name: "config",
	// 		VolumeSource: corev1.VolumeSource{
	// 			Secret: &corev1.SecretVolumeSource{
	// 				SecretName: r.Vector.QualifiedName(agentSecretConfigName),
	// 				Items: []corev1.KeyToPath{
	// 					{
	// 						Key:  BaseConfigName,
	// 						Path: BaseConfigName,
	// 					},
	// 				},
	// 			},
	// 		},
	// 	}
	// 	v = append(v, volume)
	// } else {
	// 	v = append(v, corev1.Volume{
	// 		Name: "config",
	// 		VolumeSource: corev1.VolumeSource{
	// 			Secret: &corev1.SecretVolumeSource{
	// 				SecretName: r.Vector.Spec.AgentSpec.CustomConfigSecret,
	// 			},
	// 		},
	// 	})
	// }
	return
}
