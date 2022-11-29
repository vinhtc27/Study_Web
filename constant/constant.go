package constant

const (
	ACCOUNT_STATUS_ACTIVATED   = "activated"
	ACCOUNT_STATUS_INACTIVATED = "inactivated"
	ACCOUNT_STATUS_DEACTIVATED = "deactivated"

	ACCOUNT_TYPE_NORMAL = "normal"
	ACCOUNT_TYPE_ADMIN  = "admin"
	ACCOUNT_TYPE_MOD    = "mod"
)

const (
	CLIENT_BASE_URL      = "https://example"
	CONFIRM_EMAIL_PATH   = "/confirm-email"
	FORGOT_PASSWORD_PATH = "/forgot-password"
)

const (
	CHANNEL_ROLE_HOST   = "host"
	CHANNEL_ROLE_MEMBER = "member"
)

var DEFAULT_USER_AVATAR_LIST = []string{
	"https://img.seadn.io/files/f9fe53cdcd7cd1190ec57d20fd644c63.png",
	"https://img.seadn.io/files/2f1ba270a24873cb6288041543ced623.png",
	"https://img.seadn.io/files/7c5661d4a7ed4da6360b6caeae2287fb.png",
	"https://img.seadn.io/files/49412061109d4b860a682d905ba9accc.png",
	"https://img.seadn.io/files/52f9287002753552fe8bc01b27ff7d4e.png",
	"https://img.seadn.io/files/3fb0c799e4b5673a09d2bc63e651d2b2.png",
	"https://img.seadn.io/files/f6ca3908836e94186f5148988a4acadc.png",
	"https://img.seadn.io/files/f34fa23f9c5ce63af5ea39d0f43e078c.png",
	"https://img.seadn.io/files/f34fa23f9c5ce63af5ea39d0f43e078c.png",
	"https://img.seadn.io/files/3042e1b1198f3319de82f6b44e3352df.png",
	"https://img.seadn.io/files/52450f13bc7f439e7dd9ca516068f57e.png",
	"https://img.seadn.io/files/9138c985bf1aefaca6964b77652998bf.png",
	"https://img.seadn.io/files/73712b47114c565b235bdaae4e933ddd.png",
	"https://img.seadn.io/files/94aba69a4c501f1cf2ae8ebd473f8da8.png",
	"https://img.seadn.io/files/e52f773e06875799d22df815799460e9.png",
	"https://img.seadn.io/files/c3f2b9fc8a1b93ce0c58463cc9ebba0f.png",
	"https://img.seadn.io/files/9080fe8c52155e77d7ea6c6ce97b93e2.png",
	"https://img.seadn.io/files/57f0ec32212ee4a73f4691bdeffecef3.png",
	"https://img.seadn.io/files/fa9adb2d309fa76a5a6ba6a4352a2c7d.png",
	"https://img.seadn.io/files/fe91ffffa60361aefb155b384ac95b7c.png",
	"https://img.seadn.io/files/72d5210c92b56ea8ba7ac82541851fe2.png",
	"https://img.seadn.io/files/37d9971ad38fc91be4423f58ed3244d0.png",
	"https://img.seadn.io/files/69f3472056c1d42fd544d172f019399a.png",
	"https://img.seadn.io/files/7c21b04a0b0e797faea838d7464f6a2e.png",
	"https://img.seadn.io/files/936b768d851787c0ca8bdb0de2b9bf1f.png",
	"https://img.seadn.io/files/4c5f9976fcb222065fb052953daa95e1.png",
	"https://img.seadn.io/files/b171ba566d4816dfa1706178b03f3518.png",
	"https://img.seadn.io/files/adf7cc6fc9baaffaa17f73a20fbcc9fe.png",
	"https://img.seadn.io/files/6edeca11b91f54c1d2279207109c1f3b.png",
	"https://img.seadn.io/files/2e498ab7b5bc3ca8e53ec269a9352fdc.png",
	"https://img.seadn.io/files/dde2e81c140196dc7b1d08358f0c9352.png",
	"https://img.seadn.io/files/018e4bf85a0d2040020e0165083a21e5.png",
	"https://img.seadn.io/files/0565d54fc67ef805432fa188ea644605.png",
	"https://img.seadn.io/files/fb114b276fee2b1f244dc5e2faa025e3.png",
	"https://img.seadn.io/files/a2aec0af5d5d81d138f7be11bce11f05.png",
	"https://img.seadn.io/files/079b31e8375bef3c68d6870e31b7daaa.png",
	"https://img.seadn.io/files/2459c64c944c29323c7ba3dd41f792b6.png",
	"https://img.seadn.io/files/30d46195356b95992e905696ce816107.png",
	"https://img.seadn.io/files/477e23bf79034d58a666d3cb3589294f.png",
	"https://img.seadn.io/files/a1c5971f9aac4ea32ecc2cc1e8f7ec62.png",
	"https://img.seadn.io/files/ee2b0463339dc973cffedfb1c8888703.png",
	"https://img.seadn.io/files/4e3f2c4e71ba0e19c925d5ea95d4db56.png",
	"https://img.seadn.io/files/6c538d45acdc905a17c633f9c7ce7898.png",
	"https://img.seadn.io/files/f908a6ee0aea2baa077f75a6cbfb36f8.png",
	"https://img.seadn.io/files/1b5ca7983888779a5072bdd4e4180dfa.png",
	"https://img.seadn.io/files/3a4975fc22b0e3d4a3c5f82992f0f409.png",
	"https://img.seadn.io/files/498c7a7acba360ca72b592a82fd11a78.png",
	"https://img.seadn.io/files/36d368060723ce6b88b9e8314bd33a7b.png",
	"https://img.seadn.io/files/13e983429d87386df4b4673fc942ea57.png",
	"https://img.seadn.io/files/e2b8ccb064b2c5bd5501afa8c216c905.png",
}

var DEFAULT_CHANNEL_AVATAR_LIST = []string{
	"https://img.seadn.io/files/77a76d20144a88d5989aa268a96d5e3c.jpg",
	"https://img.seadn.io/files/5620787e425bd3f6df5eb2b5d9088bc3.jpg",
	"https://img.seadn.io/files/8d84781ee3b929d44a4612a24cc80b71.jpg",
	"https://img.seadn.io/files/629adbd91faaa3df948cf71b1a1abde4.jpg",
	"https://img.seadn.io/files/eabde6e2401c6058e22525f69fabf4df.jpg",
	"https://img.seadn.io/files/e7438c722e3ac652f5b338bc5266b8d0.jpg",
	"https://img.seadn.io/files/8334752c7ccb4b1ccea5447f5098cd8a.jpg",
	"https://img.seadn.io/files/31073ba2b0723c8c2d2e25f5c68ce0c0.jpg",
	"https://img.seadn.io/files/94280918669cb0d92c895400ff79aee2.jpg",
	"https://img.seadn.io/files/6deb0d079f1d40c08591753fd66c02e5.jpg",
	"https://img.seadn.io/files/9cffc83947a3937b630d39266f6afbb0.jpg",
	"https://img.seadn.io/files/324b4a224a507e07581959d43de1f756.jpg",
	"https://img.seadn.io/files/3045c6dbbc59378b4a4322f1e556c16c.jpg",
	"https://img.seadn.io/files/adb06c1c82505977ac3f0056be5e7f6c.jpg",
	"https://img.seadn.io/files/9238048e52a3f21b50373127efd76254.jpg",
	"https://img.seadn.io/files/0d0b887323842db24953564c7d5d3f20.jpg",
	"https://img.seadn.io/files/fd52dbd97e95951ca87c54fff842edd5.jpg",
	"https://img.seadn.io/files/7bf9441f727b8affc75f80cb8dae8eaa.jpg",
	"https://img.seadn.io/files/41a843224063220e6b57ca93ce374b37.jpg",
	"https://img.seadn.io/files/8c8001996ffc9e05d344ff8a8d86b9a4.jpg",
	"https://img.seadn.io/files/4037da949b5326ce6c8db00156b37708.jpg",
	"https://img.seadn.io/files/737e84e790cb950ee87a1b5bbe09e895.jpg",
	"https://img.seadn.io/files/9e8a13cfdba37cf5926bdf4361ac234e.jpg",
	"https://img.seadn.io/files/c9e9a1a92a06e108246a2e6c83fe3fbf.jpg",
	"https://img.seadn.io/files/9f0b7979e9ac121356736802470089d6.jpg",
}
