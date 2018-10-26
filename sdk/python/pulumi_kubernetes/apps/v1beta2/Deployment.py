import pulumi
import pulumi.runtime

class Deployment(pulumi.CustomResource):
    """
    DEPRECATED - This group version of Deployment is deprecated by apps/v1/Deployment. See the
    release notes for more information. Deployment enables declarative updates for Pods and
    ReplicaSets.
    """
    def __init__(self, __name__, __opts__=None, metadata=None, spec=None, status=None):
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['apiVersion'] = 'apps/v1beta2'
        self.apiVersion = 'apps/v1beta2'

        __props__['kind'] = 'Deployment'
        self.kind = 'Deployment'

        if metadata and not isinstance(metadata, dict):
            raise TypeError('Expected property aliases to be a dict')
        self.metadata = metadata
        """
        Standard object metadata.
        """
        __props__['metadata'] = metadata

        if spec and not isinstance(spec, dict):
            raise TypeError('Expected property aliases to be a dict')
        self.spec = spec
        """
        Specification of the desired behavior of the Deployment.
        """
        __props__['spec'] = spec

        if status and not isinstance(status, dict):
            raise TypeError('Expected property aliases to be a dict')
        self.status = status
        """
        Most recently observed status of the Deployment.
        """
        __props__['status'] = status

        super(Deployment, self).__init__(
            "kubernetes:apps/v1beta2:Deployment",
            __name__,
            __props__,
            __opts__)