# Usage

[Helm](https://helm.sh) must be installed to use the charts.  Please refer to
Helm's [documentation](https://helm.sh/docs) to get started.

Once Helm has been set up correctly, add the repo as follows:

    helm repo add apiping https://udhos.github.io/apiping

Update files from repo:

    helm repo update

Search apiping:

    helm search repo apiping -l --version ">=0.0.0"
    NAME           	CHART VERSION	APP VERSION	DESCRIPTION                                
    apiping/apiping	0.1.0        	0.0.0      	Install apiping helm chart into kubernetes.

To install the charts:

    helm install my-apiping apiping/apiping
    #            ^          ^       ^
    #            |          |        \_______ chart
    #            |          |
    #            |           \_______________ repo
    #            |
    #             \__________________________ release (chart instance installed in cluster)

To uninstall the charts:

    helm uninstall my-apiping

# Source

<https://github.com/udhos/apiping>
