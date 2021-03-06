// Copyright (c)  WSO2 Inc. (http://www.wso2.org) All Rights Reserved.
//
// WSO2 Inc. licenses this file to you under the Apache License,
// Version 2.0 (the "License"); you may not use this file except
// in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
package api

const (
	ConfigJsonVolume       = "config-json"
	mgwDockerFile          = "dockerfile-conf"
	interceptorsVolume     = "intcpt-vol"
	javaInterceptorsVolume = "java-intcpt-vol"

	swaggerLocation                = "/usr/wso2/swagger/project-%v/"
	dockerFileLocation             = "/usr/wso2/dockerfile/"
	dockerFile                     = "dockerfile"
	dockerFileTemplate             = "dockerfile-template"
	policyyamlFile                 = "policy-file"
	policyyamlLocation             = "/usr/wso2/policy/"
	mgwConfFile                    = "conf-file"
	mgwConfLocation                = "/usr/wso2/mgwconf/"
	analyticsCertFile              = "analytics-cert"
	analyticsCertLocation          = "/usr/wso2/analyticssecret/"
	analyticsVolumeName            = "analytics-volume-storage"
	analyticsVolumeLocation        = "/home/ballerina/wso2/api-usage-data/"
	analyticsConfName              = "analytics-config"
	interceptorsVolumeLocation     = "usr/wso2/interceptors/project-%v/"
	javaInterceptorsVolumeLocation = "usr/wso2/libs/project-%v/"

	mgwToolkitImgConst   = "mgwToolkitImg"
	mgwRuntimeImgConst   = "mgwRuntimeImg"
	kanikoImgConst       = "kanikoImg"
	kanikoArguments      = "kanikoArguments"
	registryTypeConst    = "registryType"
	repositoryNameConst  = "repositoryName"
	wso2NameSpaceConst   = "wso2-system"
	policyConfigmap      = "policy-configmap"
	mgwConfSecretConst   = "mgw-conf"
	mgwConfConst         = "micro-gw.conf"
	controllerConfName   = "controller-config"
	ingressConfigs       = "ingress-configs"
	dockerRegConfigs     = "docker-registry-config"
	ingressProperties    = "ingress.properties"
	policyFileConst      = "policies.yaml"
	ingressResourceName  = "ingressResourceName"
	ingressTransportMode = "ingressTransportMode"
	ingressHostName      = "ingressHostName"
	tlsSecretName        = "tlsSecretName"

	usernameConst  = "username"
	passwordConst  = "password"
	certConst      = "cert_security"
	analyticsAlias = "wso2analytics"

	defaultSecurity                     = "default-security-jwt"
	endpointExtension                   = "x-wso2-production-endpoints"
	apiBasePathExtention                = "x-wso2-basePath"
	deploymentMode                      = "x-wso2-mode"
	securityExtension                   = "security"
	certPath                            = "/usr/wso2/certs/"
	mgwConfMustache                     = "mgw-conf-mustache"
	mgwConfGoTmpl                       = "mgwConf.gotmpl"
	certConfig                          = "apim-certs"
	encodedTrustsorePassword            = "YmFsbGVyaW5h"
	truststoreSecretName                = "truststorepass"
	truststoreSecretData                = "password"
	apimConfName                        = "apim-config"
	enabledGlobalTMEventPublishingConst = "enabledGlobalTMEventPublishing"
	jmsConnectionProviderConst          = "jmsConnectionProvider"
	throttleEndpointConst               = "throttleEndpoint"
	logLevelConst                       = "logLevel"
	httpConst                           = "http"
	httpsConst                          = "https"
	httpPortConst                       = "httpPort"
	httpsPortConst                      = "httpsPort"
	enableRealtimeMessageRetrievalConst = "enableRealtimeMessageRetrieval"
	enableRequestValidationConst        = "enableRequestValidation"
	enableResponseValidationConst       = "enableResponseValidation"
	analyticsEnabledConst               = "analyticsEnabled"
	uploadingTimeSpanInMillisConst      = "uploadingTimeSpanInMillis"
	rotatingPeriodConst                 = "rotatingPeriod"
	uploadFilesConst                    = "uploadFiles"
	verifyHostnameConst                 = "verifyHostname"
	hostnameConst                       = "hostname"
	portConst                           = "port"
	analyticsSecretConst                = "analyticsSecret"
	operatorModeConst                   = "operatorMode"
	ingressMode                         = "Ingress"
	clusterIPMode                       = "ClusterIP"

	authorizationUrl        = "https://example.com/oauth/authorize"
	tokenUrl                = "https://example.com/oauth/token"
	oauthSecurityType       = "oauth2"
	basicSecurityType       = "http"
	basicSecurityAndScheme  = "basic"
	securitySchemeExtension = "securitySchemes"
	securityJWT             = "JWT"
	securityOauth           = "Oauth"
	certAlias               = "alias"
	pathsExtension          = "paths"
	sidecar                 = "sidecar"
	privateJet              = "privateJet"
	verifyHostNameVal       = "false"
	apiCrdDefaultVersion    = "v1.0.0"

	hpaMaxReplicas                    = "hpaMaxReplicas"
	hpaTargetAverageUtilizationCPU    = "hpaTargetAverageUtilizationCPU"
	readinessProbeInitialDelaySeconds = "readinessProbeInitialDelaySeconds"
	readinessProbePeriodSeconds       = "readinessProbePeriodSeconds"
	livenessProbeInitialDelaySeconds  = "livenessProbeInitialDelaySeconds"
	livenessProbePeriodSeconds        = "livenessProbePeriodSeconds"

	resourceRequestCPU                 = "resourceRequestCPU"
	resourceRequestMemory              = "resourceRequestMemory"
	resourceLimitCPU                   = "resourceLimitCPU"
	resourceLimitMemory                = "resourceLimitMemory"
	generatekubernbetesartifactsformgw = "generatekubernbetesartifactsformgw"

	httpPortValConst  = 9090
	httpsPortValConst = 9095

	deploymentKind		= "Deployment"
	serviceKind 	        = "Service"
	apiVersion		= "apps/v1"
)
