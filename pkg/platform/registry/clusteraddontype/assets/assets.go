package assets

import (
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type staticFilesFile struct {
	data  string
	mime  string
	mtime time.Time
	// size is the size before compression. If 0, it means the data is uncompressed
	size int
	// hash is a sha256 hash of the file contents. Used for the Etag, and useful for caching
	hash string
}

var staticFiles = map[string]*staticFilesFile{
	"CSIOperator.md": {
		data:  "",
		hash:  "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
		mime:  "",
		mtime: time.Unix(1570537390, 0),
		size:  0,
	},
	"CronHPA.md": {
		data:  "",
		hash:  "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
		mime:  "",
		mtime: time.Unix(1570537390, 0),
		size:  0,
	},
	"GPUManager.md": {
		data:  "\x1f\x8b\b\x00\x00\x00\x00\x00\x02\xff\xbcV_S\x1aY\x16\u007f\xefOq\xb7|\x99\xa4\x00\x01\xff$\xf1-\xb3\x99\xa2\xb6R\xb3\xcb\u058c\xfb2\xb5U2͍CE\x1a\x87\x86l\xa5\xca\aT@\x10\x11fF\xa2\"QI\xfcCtlp\x8c\xd8t\v~\x98\xf4\xb9\xf7\xf2\xe4W\xd8\xea\xbeM\xab\x89\x0f\xbb/C\x95U\xed\xbd\xf7\x9c\xf3;\xbfs\xce\xefޡ!\x14\bN\xba\xbf\rI\xa1i\x1cg͏dcU\x10\x86\x86\x86\x10\xd5ӆ\xde6\xf4\x02Ջ\x82\x10\bN\"\xfb\x10)\x95\x8dޖ\xa1\xa6\f\xf5\xe8\xe9̌;\"\xb9\xff!aZM\a\x82\x93T\xa9\xd3r\x166\x1b.\x04;\x9a\xa1\xad>O\xfe\x88\xe3\x12N\xc8\xe8\x19~\x15\x11qp&9\x1d\x91H\xe9WCo\xd33\x9d\xea;\xa0l\xd3Ֆ\v\xb1\xe6\xbecm\x87в\x90\xcb\xf63E\xe8\xb4!sjhǁ\xe0\xa4\xcb\xc4KV\x96\xc8\xee\x12\xd9\xd9g\xcdw.\x04J\a6\x1b\xec\xaa\xcc\xea+\x90/\xdaH\xb6TȜ1\xa5\a{K\xb0\x94\x85\xbd%z\x92\x87\xe5\x1d\xb6\xd8u!\xb2\xd6$+\xf3t\xadAr\x17Pk\f@b\xb9\xbf\x95\xa5\xbd=C=1\xbaWt\xad\x11\bNr\x0f\x1eAp\xc0\xc1J\x06\xca\xc7p\xb0`\xa8\x05\xeepB\x10\xdc\xe8\xe1CR\xf8\x8d\xe4\u007f\xe1\x90\x1f>\xbc\xbe\xac\xf248\x11&\xea[\xfb\x03$\xd0\xfd͆a-\xb3v\x86]-\xf9\xe0r7\x10\x9c\x84b\x9dVӴ\x9a\x06m\x8d\xae\x99\x94\x96\x9a\x86\xbe\xdfO\xe5I\xe1\x03wƮ\xb6\xc9\xea>\xa9\xa5\xe0\xea\x98V\xd3\xe4M\a.Kܕ\x83\x9d\x833\xfdY\x1cZ\xc8 3\xcf\x14\x95G6Qj{\xd0*\x19ڪ\x19\x19\x8auv\x9e&Z\x99\a7t\x1d\x96\xeb.\x04\x9d6\xcf矱\uf32b\xb7\xac9\xcf\x1dspNQL\x14V]\xac8vn\x16n\xa6\\\xf5\xd7\x15C\xd5IM3]\x1f7\xa1\xf4\xfe\xab\xbe\xbe\xc1\x94=C\xd5\xc6\xc6\x1f=~@\xab\xe9\xe1(N\xc4#\xa2\xcc.\x9a\xd0K\xbb\xb8\xad\xa1j\xc1x,\x8a\x13?\xe1\xa4\xcca܉\xb8\xd6\xeeoe\a\x9c\x0e'\xe5\xd04\xe6\xf6\xdc\xdc.D\xa9ɚ:I\x1d\x9a\x94Z=C\x97ې=\xe7xy\xd7\xf7\x17\x1b\xb4\xfb\a\xd4\x1a\xbc\x13 \x9by\xe94\a4;\xec\xb4.\b\xb7w\xf9\xf9[#\x84\x9e\x86\xc3\ue604\\\bZ\xd9/OZ\xb9\x14\xbe\xf49\x87>_\x83r\x91\x1e\xb6\x90\xfd\x9bC\xf4T\x87\xed\x02\xbaY\x00]c\x8a\xd2\u007f\x97\xa6\x95M^04\x87H>\x05\xa7\xdb\u007f\x0fE\xb1<\x1b\x12\xb1\x8c\xe6\x849\xe4\xbe\xef\x87n\xaf;\xff\xa0\xbb\xa7-\xf3\xe9٤;ʳs\x87C8\x1a\x93d\x9c@s\xe8\x99\xf5\xfd\x1dN\x98hH\xb3Ė\x17\xe8B\xc7Gvտ\x9as\xea\vD \x9b\x81\x93\rēs˯\xe5\x04\x8e\"\xc7\xe7\xcf\xc9X\"\xe4\x0e\x85\xa3\x11Y\x8e\xc4$+\xa9gxv&\xf6:\x8a%3\xc0=\xae\xec\xdc?\xf3'ܕ1>\xbcP\xd3\xc8f\xd3*\xd7=\x03\xce\xe5\u009c\n\xab\u007f\xc9z{0_\xec\xf0=dϞ\xfe\x8d)'T_4U\xc3rd\xa8't\xed\x8c5/h\xe1wz\\\x80b\xbd\x9f\xda%\xb92\xa7\x9e\x9c\x1f\xb1\x8f\x1d\xb3\x15\x173\x90=\xbf\xbe\\a\xca\a\xa6ԩ\xb2\xce\x0f@o\x1dr-\xc8d \x97\x85\xdc\a\xba\xd6\xf8\x94Z\xf8\x1cx\u007f\xb3\f\xb96y[7\xf4\xb6 \xf8<\x96(Z:|WM\xf1]9\x1dH(\x94\x8e\xd8b\xd7\x16\xc2Z\x83\xab\xdc탴\x9a\xfe\xfe\xf97\xb4\x9a\xf6y|ޛ\x86\xa3\xf9\x1c\xa9\xfdnt\n\x86\xbal\xa1\xf2{\xccz:*d\xa8)Ȝ\x92Z\xde\xe7\xf5\x1a\xea\x11\x14+FהY\x9e\x99\v\x19z\x86\xc7\xf2\xba}\xe6l\xb5J\xa4҂b\xdde\xe8\xfbPZ\xb6\xd6RERi\x99\xfcT>\xf2M\xa6\xf4hW\xf1 \xb2у\x93\r\x9b\xc5\r\xb3\x00\xfe\xb1\xf1o#_\x1b\xaaf\xeaZ\xabdj\x90u\xd98q\xb9\xc0qC\v\xef\x88\a9\x82\xed\f\";\x98'\xa7\v\xce\xfcq\xd96U\xa3\xa6\xc1v\x817\xeb=5\xe0\x8eț\x0e9\xabpM\x00%\xcf\xdeg\xacz\xd0M\x1d\xba\x95\x1f\xb8z\x90Z\x11\x96\xebd\xf5\x10rm(\xb5\xfe\xfd\xd5O\x89Ĭ<1<,\xc6$96\x83=?\x8b3\xb1d\xd8#ƢÉ\x97\xd8\xff`\xc0-\xd4\x1apq`\xf4\x0e\xa1y\xc9rGd\xb7d\xa8'ח+P\xac\xc0\x92\xfe)U&\xf9\x0fpZ\xe1\x95\xff\x94\xfa\xc5즫-\xc8\xec\xdf^\xe7\xf7Q\xbf~\xde\u007f\xfbn@\x02\xbf\x1a\xfa\xb5\x14;\x98\xe7\xa8\xf9e\xe8\xf4\x97E\xc6\xf5\xe5\n]\xe8ؑ\u07b4@\xd7x\f8X\x80\xad\xde\xf5eU@\b\x99\u007f\u007f\xf9\xe1&\xa7h(\"\xd9\tE\xa2\xd3VN\xf1\xd0\u007f\x86\xf1\xe3q\xbf\xffɈ\x17\x8f\x8d>\t?\xc6\xe2\xe8\xd8#q\xcc\xe7\xf3\x8e\xfa\xfd~<\xe6\x1b\xf5\xccJ\xd3\x0fl\x1as[\xa0kT\xcf\xd2?~\x05\xed\x80#\x83\x8b}\xa3[c\x1fwX\xb7k\x1d\x1bB\xdf?\xff\xc6\xe1\x94\xdb\xfc\x99\xd4s\x8a8!w贐8zA\xabi\x87L\xbeu;\x93A=\xbe\xdc\xe1]\x1f\bN\xf2Q\xbfú0\xea\xf9_(\x17\xbd\xe3\x8f\x1f\xf9\xc3a\xd1\xfb\"\xf4\xe2\xc7\x17O\xfc#\xa3c\xa27\xfc\xe4\x85\u007f\x1c\x8fb1|C\xf9\x10z\x1d\x8a\xce\f8\xe4\xf7\xbc9\x82\xebmHm\xdai\xd5\x1a\xe6\x19C\xd5\xecW\xd4\x00\x9f}\xf9[\xc30\x98s\xb2\xab\xf2O\xc76\x8e\xe5X2.bC]\x86\xfa1d7\xa7\x12X\x12\xb1\x94\xb0о\x12\x93\xe1\x90[\x8c\xc5\xf1\x94\xebΜ\xff_\xf6Q\x1c\x8d\xc5_O\xb9\xcc\xe7\x06\xc7c\xbfRh5\x1d\x1c\x1d\xbcq\xa6\xa6\xa6\x04!4\x1b\xf9\x17\x8e\x9b\x97\xc9\x04z\xe5\x13\x84\x97\x11)<\x81\x82\xb1\xb0 x<\x1eA\x90g\xb18!\bbLJ\x84\"\x12\x8e\xcb\xd6\xcbM\nE\xf1\x84y#\t\xc2\x00\x8f\xb9q\u007f&\x13\xc8\xe7\xf5\xf2h\x038^\xcf\b\a\xe4Bc\x81\xc8\xd7<Q\xe7\xd9\xf6gB\x1b\xf1\u07b7\xc7\t\x9c@~\x8e\xfb\xbf\x01\x00\x00\xff\xff\x9fp\xa9\x90\xf3\v\x00\x00",
		hash:  "1f2c694cb54e80cc9d0bd63ff1f1362219c003f63db94cf2c73d6ba21274faf7",
		mime:  "",
		mtime: time.Unix(1570537390, 0),
		size:  3059,
	},
	"Helm.md": {
		data:  "\x1f\x8b\b\x00\x00\x00\x00\x00\x02\xff|SMO\x13Q\x14ݿ_q\x93n4\xa1\xd3\x0f$\x10w\x86\x90\x98\x10\x8d\v\\\x19\x17\xd8>\xa5\xb1_\xe9LcH\xba\x90\xd0)-m\xa5\xc4J\xcaL\x05\x8aTf\xa13 \b\xb4\xd3\xd6?3\xf7\xbd\xe9\xaa\u007f\xc1\xcc\f`\xa5Ʒ\xbc\xe7\xbes\xcf;\xe7]\x9f\x0f\x1e\xd3x\xc26\xceY\xe3\x03!>\x9f\x0f\xb8\x99\xb7\xcc\v\xcb,s\xb3J\x88\x83\x02k\x18\\o\xf1Z\x01\x16\xb3\xafh&I%*\x02v뼮q\xad\x8c\xdd-\xae\xe4Y\xe9#Vd\xbcl\xa3|9\x05L=\xc7#\xc56\xbe\xb2u\xd96.\xd9~\x9b7\xcb.\x17\xea\r\xde\xdff;\x1b\xac\xf5\x85\x10[\x1eغau\xb7Q\xef\xe0\xaeƚU\xdcl\x8dz\xc5%\x9a\x8cФ4>o!\xf9&\x96\xa4\xa3^eiqa\xd4+\r\xd5\x02+\xd6\\\xf5\xc0\xd5+\x94\xcfps\xdf^\xef\x8fz\x15\xb6U\xb3\x06\xaa\xd5-x\xe8\xfc\xcarF\x02lj\xac\xb2\x81\xba2T\v|p\x84\x05\x19\xd5\x01\xf6\x0f\xb1\xb2Õ<\x1e\xeea\xf1\x80\xd5;l\xbf\xed\xd90\\\xd7x\xff\a6\xb5\xdb\xfe\xb7\xb7Z\xd0\xe8ا-B\xc6Q\xaf\xdf\x1d\xf8(\x1a\xf5\xa7\x920\x05xR\x98l\xb1̶uU\x9e$\xcb\xc1\xdd\x1a֪\xfc\xf8\x04r\xc0OM\xdc+\xc3\xcd\xc9\xc1\xd0l\xd8\xfa\x11V\x0fx]\xb3\u007f\xe6Y\xb7\x060\x06\xb3\xd2{<\xdd{\xba\x9c\xa0bz9BEȑ\x1c\xf8\xff:0^\x80\xbb\xe8D\xc9e\x10\xdf\xc5^K\xe3s\xa24\x1dO\xad&\x9c\xa0r\x10\x14\x82\xd3\xec\xe0j\xfe\xd9\xf3)\b\a\x9f\xc4\x1c\x87\xbf7\xc0{\x96_\\\x15%\x9a\x00\x97G\x8a\xc5\xe34\xe3\xf7n\xff\x83'4s\xc33\xf7\x1f\x1e'&7\xe0\xe1n\r\x8b\x17\xecs\xcb2/\b\t\t`\x992\xab\x1b\xac\xb2\x16\x12\xe6x\xa9Ț\xdf\\\xcf7\xb9\x92\xffc\xb1\a\x90\xb0\xe0\xa6\xe4:\xe9\x05\x05A!<\a\x9e\x02\b9\n\x80+y\xcf\xe5ۙV\xff\x17\xafkl\xa7\xc3\xce>\x11r\xfdϏ\xd7^8\xd8\xf5b\xb8\v\xf3\xf2ފ$\xa5Ň\x81@$\x9e\xcaF\x05\xc9\xfb\xd7B$\x95\bDS\x91\xac\xf3\xe6@:\x93\x8af#R\xe0\xc1\xccl`:<;\x1d\xbcO~\a\x00\x00\xff\xffQ\u007fq\xaf\x98\x03\x00\x00",
		hash:  "fd1482cf5a9793289e8a5fc67c408ef9e13b87cff16f40a64dc49ca02de72cd6",
		mime:  "",
		mtime: time.Unix(1570537390, 0),
		size:  920,
	},
	"IPAM.md": {
		data:  "",
		hash:  "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
		mime:  "",
		mtime: time.Unix(1570537390, 0),
		size:  0,
	},
	"LBCF.md": {
		data:  "\x1f\x8b\b\x00\x00\x00\x00\x00\x02\xff\xecZ\xebS\xdbf\xba\xff\xee\xbfB\x93\xec\x876Y㻱\x99\xe9\aI\xbe\x82e|\x03c\x9f93\xd1ͲlY\x12\x92l\xd9lv\x86&\xe4J\b\xe4L\x936\x844%\r\r\xdb6@\xb69Y\xb0\x81\xfc1+\xc9\xe6S\xfe\x853\xd2k\x83!d\xd3v\xbb\xdf\x0e3\x9dJ\xef\xfb\xbc\xcf\xf3\xbc\xcf\xf5\xf7ȹx\x11J h\xa4\xb7\xfd\xc6\xf8\xe6\xbe\xcdv\xf1\xe2E\xa8\xdbY\xd0:o\xb5\xceb\xb7\xb3\x04\x8dA\t\x01\xa7 \x04\xe7p\x9e\xa4%\b\x15xE\x128\x8e\xe5\x19(\"\xe15Z\x15\xa4*\xf4\x99\xc9\xe3s\x9b\xcd\xfc\x9f\xf1Ͷ\xb6;o\xfc|xt}\xb3{\xf0w}ms\xa2N\xd0\x12O+\xb4\xac\u07fc\xd1]]8\x9a_\xed~\xb5\xd9{\xf3\xacwp\xa0?\xbd\xd5[_7\xee\xbf\xd4o\xbf=z\xfa\xdcX\xbfi|\xfb\xf6\xfd\xfe=\xe3\xebM}m\xf3\xe8\xf1\x92vp_\xdf\xda\xd3\x1fo\xea\xdb{\xc6\xfd\x8d\xe1S\xdd\xd5\x05}\xeb\xdb\xee\xfd\x9d\xa3\xd5C\xbd\xfd\xc3\xfb\xfd{\xfa\xde[cyE;|\xa2\xef\xb7\xf5\x17/\xbb\xab\vƝ\xbf\xe9\xaf\x1f\xf6\xae\x1f\xe8w\x9fh\x9d\r\xa3\xb3\xde{\xfb\x8b\xb6\xbb\xaa\xdf]7\x1e\xed\xe9k\x9b\xda\xc1\xbb\xb3\xca|\xfd\xb6\xbb\xba\xa0\xed\xfeh̿\xd4\xef=:Z\x9b7^_\xfb\xe7\xfc5`\x9d\xe3[\x1d=\xb9\xd9=|\xa1\u07fcQ=\xb9\xde\xf6^\xef\xf5\xba\xcd6\xbc\v\xe8M\xc3@0E\xd9\x05\x1e\xfa3\xa4\xef\xdc\xfc\x90D\xeblh\xbb\x8b\x1f2\xbb\n\x9d]\xd3W\x96\xba/w\xa0O\xfd]\x85\xba\xaf;\xfa\xb7\x8b\x9f :\xea|\xd3\xdbz\xa1/}g\x9a\xe1\u007f\x17\x8c\xf6\nt\x152\xee\xcc믿M\xe25Z\x16q\x92\x96\xa1\xab\xb6\xab\x90\xfd7\xfdA\x9f:pLp\x86Ғ\xc5\x11d\xc9N\xf6\x83\x8d\x96>}\xd7\x10-rB\xabF\xf3\xcaG\b\x1c\x83\aӜv\xb9%+t\r\xfa}\xb2\xb2\xb4\xd4`I\x1a&I\xa1~\x8e\xbc?T\x16\xca\xd5e\x85\x962\x02G\u007f\xe2^\x8eS\xeb\xff\xa6,\x84\xe5)3\xc7\xffc\xb2\xb24)\xd1ʿ \xf8\xe3\xfd\xf5{d\x118Y\xa5yJ\xa2IA\xa2\xe4\x11S\xf4\x88R\xa5e\x05'\xab#\xacpl\xb9\xba\xac\b\xb5\f-\vu\x89\xa4Ct\x89\xe5Y\x85\x15\xf8_a\xb9\xbe\x04F\x12\xea\xe2\xc7\x04\xfc{\x128\x01\xa7\x88~\x01\xff\xcfK\xa0$\xb6q\xae\x9c\u007fO\x82\xe9\xf3Z]\xc1\x15\x1a\xfa\x15\u007fW!̤ey&O\x13eA\xa8\xa2\x02_b\x99\xba\x84\xf7\xe5}RV\x03\xe7X\xeaWI\xbb\nM\x03ڏI\xfb\xb8,\xab\xa3\x98\xbd\x01t!}\xadm<\xde\x06mT\xdfޛ\bdAw0\x1e\xff`\xbc\xfd[wu\xa1\xf7n\xa5\xb7~\xcfXk\xeb\xb7\xdf\xf6\xde=\xe9\xad\xdf\xd3\xda7\xf5\x9d/{\xdf\xdf\xd0\xf7\xdej\x9d\x8d\xbe|\xb33\x1e<\xd7\xf7\x97\xf5\xed=\xfd\xc5#c\xf5\xcd\xd1\xda\xdf\xcd\xfe\xb8\xb6\t\x9a\xa9Ipc\xb3\xfb\xd53\xfd\xc1\x81\xfe`\xd3X{\xa6\xed\xbe\x02\xad\xd3\xe4\xf8b\xb5wx\x18\xe8\xbe\\\xea\xf3\xfb\xe7\xfc\xb5\xa3\xf9\xd5\u07bb[\xa0\xdf\xf6\xde=\xd6\xdaO\xfa{&\xdb\xfdy}\xf9\x81\xd6n\xeb\x0f\xbeї\xb7\xb5\xceF\xef\xa0c<=\x04\xd4\xda\xeeb\xefݎ~\xf7Y\xef\xfa\xc1\xfb\xfdU\x9b\xcd\x0e\x81^\xaeu:\xc6\xc2\xf2p\xebu\xe8+K\xfa\xab\xaf\x8d\xb5%\xfd\xee:\xe8\xe7\xbd[?\xea[\xab\xdaޝ~\xff\u007fw\xab\xbb\xb9h\xb2\x00\x9a\fv\xbb_\xee\xe8\xed\x1f\xf4\xa5[\xdd\xf6\xcb\xee\xabG݇\x1b\x16\x8du\xd7\xfb\xdb\xfa\xf3\xeb\xda\xee}\xfd\xc6[\xad\xf3\b\xbc\xea7^k\xed\x9f\xf4\x95{\xda\uef36\xfb\xe3\xb0\x126;\xd4{\xf9}o\xfe\x87S\xa0\xe0Ꮁ\xb4u\xf4\xf4\xb9\xb6{\xff\x18\xad\x9c\xf8\xef\xe8\xf1\x8a~\xfb\xad\xf1t]뼵V/B\xdd_:\xddγ\xde\x0f_\x1a\xaf\xaf\xf5\xaf=\x11\xc8B\xae\x11\x97S_\xbeku\xfb\xbb\xdd;\xb7\x8d\xb5\x9fMM\xf7\xe7\xf5\x95\xedP\x8b\xc7k,\t\xc1T\x8d\x95e3t\xfap\v\xb8\x0e\x17Y\x99\x96\x1a\xb4d\xfa\xea\x1f\x1d\xfd\xeew\xfaʶ~wS_\xbef<\xdc1e@\x90\x1d\xb2\xdbi\x1e'8ڎ\x0f\xb8\xd8E\xaeΰ\xbc\xfc\xc5 '\x8e\xf9\xf7]\xf8\xe7\x93\x00>\xbb5\xa45P\xf6CU\x8e\x9e?5c\f(tF\x95\x12\x8d+u\x89\xb63\xb8B\xcb_\x9c\xce\xffl\x9d\x90\xfa\x8f\xf2\x17\x8aT\xa7m6\xe3\xfefoi\x05xȲ\x99\xbe\xb6\xf9_\xbd\x1b\x87\xbd\xadm\xad\xfd\xe0\xbf?++\x8a(\x8f9\x1c$'ԩ\x11\x85\xe6I\x9aWFH\xa1\xe6\x10%\x81\xaa\x93\x8aC\xa9ҟk\xbbw{o^i{;\xae\x11\x97{\xc4\v\xf4\x06h\xcbB\x96\xdf\x1d\xad\xcdkﶌ\xaf\xf6\xb4NG;x8Pڌ\xe3\xee\xd6k}y\xbb\xfbզ\xe9š\xd44\x1e\xed\x19\xbf<\xb4\xd9\\#\x10\xc8\x04\x00,\x01Vַ\xee\xf4\xbe\xbfa\x12\xdbl\xee\x11\b\xa4\x83q\xfb\xd1\xd1\xfc\x1dc\xf1o'\xbbPߨ\xbd\x97\v\xbd{\xd7\xcd|\xb6\xa2\x03\x84r\u007f\x0fľ\xfex\xd3f\xf3\x8c@\x00\x14\x1a\xf7\xee\x1c\x1b\x01M \x80\u007fwuA=sD\xdbmk\x87\x8b\xfd\xf8\xb3\x04\xa2\t\x04\x02\x05\x19\xac^\x84@\x1e귿\xee\xado\x9a1ٯ<\xff\xf8\xbb\xb1v\xe7L\x16跟蝶\xf1h\xc7\xd4th\xeb\xfd\xfem\xfd\xc9\x13\xfd\xf55\x87\xb6{]7\xa3\xfb\x8e\xcd\x0eu;\x0f\xf4\xad\xd5A\xabM\n\x14\x9d\x12$\xc5f7u\xe8>ys4\xbf\x9a\x9a\f}\xd6}\xf2Ƹ\xbf\x01hS\x02ջ\xf5\v\x9a@\xde\xef\xdf\xd3v\x97\x80]\xfb\f\x00O\xe3\xe9\xf5\xa3[K\xbd\x9d\xeb\xc6\xc37fZ^?\xd0_<3\xbe[?\xfaт\xfa\x8b\xff\xd3\xed<=z\xfa\xad\xf1K߇\x83K\x0e\x81\xeea\v\xf4\xb1\xfb\x9d\xa5\xe3\x00(\xc0X\xe2\xfd\xfem\xe3\xd1-3\x90\x1f/\xe8\a\x0fMSX\x85\xb0\x85\xd78s\xa3c\xce#Gk\xf3\xbd\x1f\xbe\xd4W\x1eP\x16\xce\x1c17\xb5\xddW\xfa\xfaO\xfa\x8d\r\xe0%\xedݺ\xf1\xe5\xb6iӣ\x87k\xfa\xf5\xe5\xfe\xbb\xdd\x02\xd3k\x9b\xfaڎ\xfe\xec\xd9\xf1kC$\xe3!\xe8\xfd\xfem`\v\xf9\x8c\xdd\xcc!\xe4\xabM\xe3\xe9\x86\xf1løsػ{\xad{mϬ\xe6\xed\xaf\xba\xab\v,/+f\xb3\x8d\x87\x80\x9dd\vL\xd9Y\xea\xe4\xb9J\xb7l\xb6+W\xae\xd8\xccN#\x8b49f\x03]G\x94XAb\x95\x16\xca\xe1\xb2l\"\xfc1\xe8\x02@<v^\xa0h;)\xb1\nK\xe2܅>\xbd\t\xb4p\x96\xa7%y\xc0\xc1Lm\xde:\xd87\xecIKck8C\x8fA\u007f\xfa\x8b\xf5`7\xa9\xfe:\xb4\x8dK\xcc\x10\x17\xc0\xe9\x82\xdd.\xd1\f+\xf0_\xfc\xe9/-\xa1.\xf5\xdf\xfez\xe1C\xba\x86H\xdaYj@\a\xdeΣ;6ǀ\x14\aH\xfdd\xe3_\x9c\xaaҭ\x8f\x1c\xabҭ\xbf^\xb0lj\xeb>\xee\x1c=\xbey\\S@\x12\x810\xd0\x1f\x1ch\x9d\x17 \xe9\xcd\xf0\x02^0\x81%\xa9p\x10.\x8a\\\v\xb2\x97L\xb3\x96X\xa6\x86\x8bV(}\xb8?\x14g\x1fn\xf6\xa3\x05\xecZ\n\x81\xb8\azt_\x802p\xe9\xd2pr\x83\x94E\x13ȥK6\x9b\xb1\xf6\xb3v\xb8\xa8\xed\xbe\xeak\u07be\xc9R\xdan\xfb\nG\xd8GՒ'\xe8\x95\x1aW\xced\xbd\xbe\xf5\xadv\xb8\xf8~\xff^\xf7\xc9\x03}\xe5gPn\x00\xd3\xe3\x15s\xf7\xa7m}\xf9{}\xf9\x1f\xdan\xdb\xedt:\x9df]]Z\xeemm\x81i4\x87\xa6\xacY\xfa\x92\xf1\xcb\xe6\x18\xd4\xdd\\\xd4\xdb\xcb\xda\xfe\xaa\xd6ٸr|\xf4#\xe7\xaeh\xbbm\xd0b\x8dg\x1b\xbd\xed\xe7\xc3b{\x8b\x1b\xda\xee\x92\xfe\xea\x1b}\xcd|\xd5\xf6WMhpws\xb8\x86]\x02\xae\xc0Ev\x9a\x96\xcc\x1e7\x06\x9d\x05\xa9\x8e\x86\x8b\xa0\x15\xdce\xab\xb2<\x05\xbe}\f>}\xd8j\xb4\x82S\xb8\x82\x9b\x01\f\u009fn\xe25\x91\xa3\xedB\xc9N7Y\xd9\xec\x9fv\x8e\x80\xfa\xfb\xd6\xe4<6<R\xd8\x06y\xc8\x11!+s\xc6\xfa\xb3\fG؏S\x89#\xb2\xc7\xd9\xca\rɏ\x87Ơ\vC\xfe\x01\x11̱\xb2B\xf3\xb4d\u058c1\xe8\x82e\xb93;\x92\xa0\b\xa4\xc0\x8dA\x17rh\xcaܣy\xb9.\xd1)\x81c\xc9\x16\x90#\x82g\b\xe6T\xbc%\x83\x88\xbati\xd8v\xa0ҟ\x8a\x1eP\xc4 芙\x88\x84\xafL\n͊\xf7\x8aY\x14\xads&\x8e\xbc\xf1s\xf7\xe0\xc1g\x93\xa9p\xf2\xf3sC\xc9\x04\xad\xbbmmo\xf1\xf8\xc8p\xf8\x04\x83\xc1`wu!\x96˥\x86]m\xac\xcd\xeb+\xf7\xb5\xfdU}ms(\x16\xfbb\xaf\xd4Z\n-\x03P\xc0\xf2\x14\xdd\x1c)+5\xce\n惟\xf5\xe5\a\xbd\x97\v\xfa\xed\xc7\xff\xc1H %\x1aWh;O\xab\u007fl,X\rc\f\x1a6\xf6\a1\x92k\x89f97\r~nx\x98\x16\xfdht\x98v\x06\x9b\x94P\xc3Y~\f\xbapbK\xb0Q\x97L\xc2!\xbb\xfe\x96p\xeam\x1d\xea[\xab\xfdq\x17t\xf6\xa1`z\xfd\xd0,[\xed\x9bg\x01Dwu\x01t\xc8\u007f\xce_\xeb=\xff\xe9\x14\xca0\xa7\x18\xeb[\x98\xb6ۖ\x1b\xa4\xddT\xf6\xfd\xfe\xbdA/\x15\x05I\xd1v\xdb\x01\xe7\xfb\xfd\xdb94\xf5\xb9Y\x9d\xac\xe3\xfa\xed\x1d\x13\x94\xac.\x18\xdb\xcb\xda\xee\x8fWLYc\x03\x81f\xa8\x00펮\x1f\x18\xdfl\xfb\xfd\xbf9X\x10pǨ9ҟ\x13,*M\xd8Mu\xfb\xa6\xb0[\xa3\xff\xaf\t\x14Э\xcd[Zabz\xde>\x18\xbbmР+\x00\x1f\x00Q\x03\xab\xf4\xdd\")\xc7(@\x90\x94d\xbdF\x98q\x17p\xda H\xc4%\xbcF+\xc7m^\xa5Y\xa6lƌ\xdf\u007f\xe1\xb8\xc5@\x00\"\xf5\x11\xe60&\x05ps\x00\xb4\xc0\xfc\x8b\xe1\xa2\x05\xe2\xcfZ\xaf1\xb0\xd31\xd99FR\xa4\xba\xacД\xbd\x0f\xf2-ď\x8b\xec\xc7\xcd48~\xe6\xc0\x88H\xd7Ơ\xab֝\xacϋH8\x1aOBh8\x93\x8bG\xe2(\x9c\v[\xab\xd6>\x16\x8fG\x99\x10\x8a\"y\x96\x81\xd58\x023q,.W]\xce\xcb\x05w4:35\x91\x93j\x18쌢\xd9\xd9h6NxB\xe90\x82\xaaS0\x16-4\xd19x\xdcb\x820\xc9i\x04.\xe4\xe0\xeax\x04\xcbTՈZ\bM\xa7\xd3\x13\xe1&\x12#\xa2A\xb6\x90o\xe6\xf0<U\x8fG}\xee\xc4L\xb2\x8c\xe5Tu\x92\xb1hB\xe1\xb9d\xccbr\x8a0\x1cl\x15ݑ:>#\x96\xa9(\xd7 X$_\xc87\xc5b4\xe2\xc4\xf3\xc1z<\x9cD⨓\x99r\x96\x11\xac2\xe56\x9f-&\x99\xb9\xb8\x1a+\x93I,\xc74\xb1J\\\xc5Bd3\x99\x83}yk\xadj\xae5\x8f\xd7*(R\t\x85\x13\x18\\\x8d®\xa90R\xc6P\x8bI\xda\xe9m\x86+p\x1a\\\x8bɡy\xcaU\xc8\xfb\xf8b4X/΅3\x18\x1c\x00\a\x9aX\x9ct\x979\x82\x17\x85b\xde\xdbLV`\xa7y\x06\xd8\x04\x16r\x89\xe9\fGԒ\x1c\xc1\xa7\x99\xe9\xe8t\x05\x8f\xfa\x1aD4\xc8\xd3YX\x98:9(\xc6\xc3\xc9\x061\x83\x94\t\xbe\xca䢜\x82\xcfd8\x8bI1\x14\xa61D\xb5\xa4\xc1\xaa:=\xc1\xfa\x9cż\xafR\xcc\xfb\x9c\x05w\xb3A\xe53e2Z\xad\x17\xdcA\x05\x8bǑx\x05N\"Lu\xb6\\e\xa3AՉXL\xe0t8\x02Ó(\x9c\x0e\xc0&\x11\xcaL\xa0p:\f\xb7h\x91\xa0\nN\x0e\x8f\xbb%\xa6Ts\xcd\xd2\xde\xdcT\xc3Y\xcd\x04%ga<\xad\xe6ԙ\xa4'\x81\xf9-&\xee\xfc\xa8\xafૹ\xab\xc1<Y\x8a\xb0\x1c)\xe0\xe5\xe0\xacw\xdc\xc5SUBaCY<\x91)\x8c\x17Hz\x94mL\x12\xb5\xcch\x90!\xbcX\x14\x95\xf8i\xb6\xe9\xa4\xe3\x16\x93\x89B\xde5\xcbF\xe2H%\x9ck\x96\xe7rE\xc5+O\xcf)J˛\xc8;#\x973^\x8ea$\x1cN\x05\xa8\xd1YJ&\"\xb9V\xd5\xdb\nR\x13\xb1\xac+\x02'ӤŤ9]mR\x918\x86V/\xe7\x05Q\xc4j\xb9\xb8;$&\x99`\xb5\xe0G\xa4\xc4h\x8eM\xb4*\xcd\x00#F\\3\xad\xda\x04\xe7D\xe9\xd2l\xa1\xa9\x16\x15\xc7\xe5\x890\x0e[LF}QG\xc6\xe3w\x95\xdd*\xeb\xe3\"\xa8\xabE\x94\xa3\n\x97\xca\xf9\x15w\xc0\a\xfb+~\xbc\x81\x8e³\xe12UL\xcf\np\xc3\xefd\xd3u\xca\x1d\xab\x8c\x12\xb5*Q\xb2\x98L昰\xdb{ٗ\xa0E\xba\x9aoM8\xd9p\b\x8d\xcd\x04\xf0\xa4CUJ\x13\xa9\xc4,\xec\x11\xd2\xd3\x1c\x9fHpD:\x1e\x82\xd30\"x\xe3!$\x87\xa2\xb0\x1aV-&!3\x0f2\xce\x14\x9c\x8e9\x108\x1d\x82\x99<\x83ţ\f\xc20\x12\u0084#H\x9aD\xe0t\x16\xa9ΡH5\x84$\x87\xd6UX(\x02\x17㱌\x93\f\t\x8d\x84'\xc9\x15<ӭb\xd6\xc7\xf7\xf3i\xcȩD-\xd9 \xb2\xc1J!\x9f\xe4H>\xdd(z\x92\r\xb2F\x95\x89\x1a\xe7/\xccdD\x90\x80n\x9f\xdb:\x10\r\xb7\x8as\xe3-,뫐|ZM\xa9\x85\xf8\x84Z@\x90\xf4T\f\x83\xa3Q\xcc]vR1؟h\x05\x1b\x05OR\xc5X \xccbrF\"O\xba\xcfd2_,\x13\xb1\xa4Pȍ\xf3X\x05ɛ\xf9\x12\xcb\xc2\xe1\xdc\x1c\x92\xc4\xc2\xe1\xa8\xc5\x04\xcd\xcaQ8\x9dF\x84P\x1c\x89\x84`'\x16\x8aGQ\xd4\\\x9b\x8a j\x1c\x89\xb0E\x81\x8aeTrμ2\xe5\xa1Z\xa7\xafk1\xe9k\xd0*\xce \r\xd2\xcd9\tϸ/1\a\xc7\x11\xa6\xc83iլt*\x9a\xb6\x8c\x9fC\xe0\xb8\n\x87\x90qK\x9bL \x9c\xae  N\xb0\x90\x9fIM\xc0\xfer\xa5lI\x9cd\x03\x8d\x82g\\N\xd4(\x99p\x8f\x9bW\x11\x8bn\xaf\x95\x80\t75g\xfeGx\xc6\xf9B\xde'ҵ\x88\x13x'\x1f\xacS\xb5\x88L\xba\xcbe\xacF\xb6\x12\xb5d\x8b\b\xc1. -\x1cNT`\x99)O\xcd&\xf8A\xb9\xc8T\x88h\xd0U\x8cFT<\xeb\xab\xf4\xbd\xe3FÞ\xf3\t\b\xb7S\x8d\x81\xabp\bRP#\xf0\x89\xc7`5lٮ_\nL\x03b(\x868͚B\x85\x98t\x1eA2\x19\xbf\x87$\xcas\x01\x98\x13\xc8r\xb0\x18V\xfd\x05o\xae81#\x8b\x15\xb8\x04<\x84\x85\xa3!8o1a\x90l\xdeS\x8b5\x1b\x195?1\x8d\x91*\xd6\fNz18=Y\xc0K!\x14\x85Ӆ(\xca\x02\x0f\xba\xf8*\x8a\xc0\xf10S\"\xc3L)\r\a\x18\xc4\x034\x89\xe4\x11\x17\x16\xcfWC\x05$[\x17\x125\x17\x19\x98r\x84`_\xa8\xec%ѩx\xa4u\xb92[v\xc6\xc2A\f\x83a\x18)\x06D\x8f;\xe5\x85a3E\xc2\fȝ\f\x13/ÓHV k*]\"\xbc\x1e.B\xe4\x83\xe84\x15\x9ch\xfaS~\xc1\xe3+$\x84\x99\x8c\xaf0\xe9oഛ\x80\xd90\xacNg#9\u007f3N\x8c\xaa\xc0&\xb5\xf4\xec\xb4:\x15W2S\xd1`B\x99\xad\xd4\xd3\xc9r\t\xab\x96S,:홓aJ\rճ\xa3n\x0fE\xfa\x98z=\x82\xe2\x99˒\xb7\xe8\xab\tj\xf0\xf2\xcc(\xd2\x02\xe51\x8e\xc1\xcdX}\x9c\x9e\xad\x04[*\f\xc3\xf9\xd2\x04\x15,L\xc10\x1c\x86U$\x8e\x85\vh<\x8d\x06\x94&\xefM\xb8\xb2\xb9t\x10g\xbd\xe3d~\xdc\xdfP\x93\x02\t`\xadO\x89\x94\xd2\xf2\xc4L(*{љ\x18>e^o|4\x95\xcb1x%냃\xcd\\#G\x85\x9c9\xd5\xc3\xc6\xfc^\x0fV\xa9P\x89\xdcD,\xe0ISu\xb7\x0249\xafY\xc3\u07b8YZ\"u\xf78\xd9JD=\b\x13(\x94\xc7\xd9\xcb\xf1Y!\xc7UԨ\\@\x83\x01\xd6\xe5\r\x94\x85\x1cH@\x95#Ԙg\n\x8fI).\xe3\"\xa5\x99@\xacqy&,\xc7*\xdeI\xc5\xe3\xa8J\x14[/D\xf3\xe1\xe2t\xb9\xe0wсPΓ\x16\x1aS\xb9\x99\xb2_lD/\x03\xef\x04\xd3nw\xb6\x14\xc5\xeaѺ*19\xaa\\\xca\xf8҅\x86W-\x87\x1dXś\x9b\x1dO\x872Q\x84V\x88 %\xa6\xaa\xe5r\"\xe9\x0f_.;\x82\x8e|-\xda\x02\xbd86ILM\x15[\xe1\x94K\xca\xcd2\xa9f\xa2\x1a\xa0}3-%\xd0j88A\xf0Ѣ\x92\x9eiP\xf5\x82\xa3\xe1\xe5\x19\xb11\x9e%gi14CS\xd9qʓ\x9d\x006\xa9\xd7G\x19ڄ\xe8\x14\xe3(!\\\xb1R\x12\xabr\"\x9f\x9d\x8aNOק\xa7\xb2\u07b9\x9c\x9b\xc1117\x1br$+͘\x8az$!Φ\x82b\x96/\x8c\x82`\x8b\x92\xa9|\x89\x17\xfd3r\x00/\xa35\xb6@M\x86UB\xc2b\x8e\xb4\x12\xa62\x94\xc0˩\x16\x96uG\xa6<\x99\xc6h9t\x02\xbf\xc2\xc9Ї\xe0\xcb\u0097'\xbf\xe0Z\x90\xd1\xfa\xb6\xf1;'40Q\x9d\x83*?\x1c\xb1>\x01\xba\x01\x1d\x98\xab\x06\x9f\xc4\asPYQ\xc41\x87\xe3\fϑ!6#r\x83\xbc`3o8|\x11\\\x14e\xc71\b\x1e\xba\xf6\x1f\xa0\xafD\x8b\x1cK\xe2\xf2\x18\xe4\xb2&\x02\x8e&\x15A\x02\x90\xbe\x86+d9\x81\x134w\xfc\x11\xee\x03\x9b\x92BM\x14x\x9aWΓ\xad\xd05\x91Õ\xfe\x841\xac\xad\xc5\xeb\x14\xe7\xdf\xca\xfb\xff?R\xfe\x91\x1f)ON\x99\x83\xde\a\xd796\x17\xf8.\x10p\xdaN\xff\xa4\t\f\xc7\xf22M\xd6%\xfa\xac\xe9Ru\x8eK\x9d\x1e\xf1O(\x1a\x02W\xafј\xa9\xd2\abOOv$~Fj\xcd<\x94\u0095\xf2\x18\xe4\xa0\x15\xd2!˜\x83\xa4%E>C'\xd185\xc9s-\x8b\xd5@; \xf7\x1c\xf7\x9f+\x8e\x1cL\x9e\xa7U\xfcף\xa7U\xa5\xb2\x83!\xfb\xdc\x1au\x9c\xd5}\xb2S)}\x92 \xbf3\xb9\x87|i?\xdfG\xe2i\u007f*\xb8\xc4\xd0ʐ\x93O\x17\x84ߖ\xa12m\xfd0\b\x97\xac\x9f\xed[cPR\xe0M\xb9\x8aU\x1c\xfb\xffv$\x9e\xb2\xcc\xf4\u007f\x01\x00\x00\xff\xff\x12K@[W&\x00\x00",
		hash:  "8f1304a7838321a396aa962a55d306ead51b8433dfc80af43c4b1a8df1c0c4f9",
		mime:  "",
		mtime: time.Unix(1571217148, 0),
		size:  9815,
	},
	"LogCollector.md": {
		data:  "\x1f\x8b\b\x00\x00\x00\x00\x00\x02\xff\xbc\x95KO\x1bW\x14\xc7\xf7\xf3)nŦ\x91\x82!v\x1c\x93\xeeR\xbaKZUJ\xbb\x8a\xba\x18\xc6\x03E`\x0fŎ\xaaJ^L\x8c\x9f`\xe3I\x04\xc6\xd8\x040\xc5\xe0\xa4bƔ\xc4\xd8\x1e?>\fs\uef73\xe2+T3w0c\x1e\x12\xabze\xdds\xefy\xfc\xce\xff\x9c\x19\x1bC\xaf\xa4\xb9iiqQ\x14\xa2\xd22վ\xe0\xd2:Ǎ\x8d\x8d!\xa2'\f\xbdi\xe8kD\xcfs\x1cު\xc1`\vo4\xcdJ\nV\xf7\xe8J\x0f\x974P۰]\xc7;yX\xad\x1a\xad\x0e٨\xe3\xcc9.(F\xbfB\xca\t\xb3\x92\"\xfdCH%Gޞ\xd7 y~\xd9\xcdAA3\xf4\x1a4R\xd7\xd7l?8Sd't5N\xe2m\x92m\x83Z\xa6\xe7\x1a\xf4\x13\xb8\x986\xf4&)'\x98C(\xbc7\xe5w4}\x86^\xf2\xb3\v<\xb2\f\xb94\xa8e\x14\x95\x96\xe6\x05\x843E*'\x91\x13\xddv\x8e\xa6_\xbd\xbe\xbe\xc7,FK7\x0fJ\x17r\xfc\xae\"\xcd\x1d\x99\x1e\xbd3Z\x1d\xac\x15\x8c\xd6g\x96\x19ή\xc1j\x1d\xba2(څ\x1c\xbf\xfd\x8a\x99@Y\xbf\xec\xe6\xdcV\xf4bN\fG\x91\xd1-\xc3N}X\xb7\xa1\xd7\xd0\x0f\xbc\x18\x92\xc2\x111je\a\xbd\x03\xe8\x16\xe8@\xa1\xd5܅\x1cgX\x19/S.\xd3A\xda\xed\x93\x1e' \xb3m&\U000e49f2s\v}:mVR\xb8\xa3\xc0\x87\x1cnf\xe8\x976\xf9G\xbb7\x19C_g1\x98\x17\xf7{:\xa8\xd0\xea\xc83\xabu\xed&4RN\x13RIK\x05W\xadp$`\xe3\xb58_Ŷ\xf12\x98\xf8\xac\x8e\x13\x05\xcbX\xb2R2z\x03\xeb\xd1=\xe4q\xbcN\xaa*U\x0f\xd1˷3\xe2rX\x8c\x8a\x114$\xc7$b\x01?ܣ\xea\xc0\xdcr\b\x8c\x86\xb5\xc4l\xae\xd4I\xef_7\xf5\x85\xa1?\xd0\xda\xf4\xb4\xcaqn+\xbb\xef\x1e\r\xf4\"\x18\x1c\x97\xc2\xe81\x82F\xea\xf6UC\xaf\x19\xad\xb5\xdbNc\xe8\xe6\x19(yr\xdc@1DNu\xd8]C1d\xea%\xaa\x1eB~\x9fl\xd4\xe9\xd7\x04\xee((\x86pV\x86\xd3ݟ\xf8\x90\x18Y\xe2\x051\x82b\\\f\x8d\xdf\xfc!\xfb\f\xb9-\xe8\xc65\xfb\xe1\xa247.\fK\x891\xb9\xbd\x16\xa3(\x86\xb5\x02\xc38\xe9\xf1\xe1\xfd\xd6\xf4Ͽ>F^\xff\xe4\x8f\xdf[\x9d=)Ŭ\xf4\xc7#\u007fE\xa2b(f\xb1\x1c\xd9\x17\xacy\xb0\xd3\xc1\xdbڝ\xe3c\xab\xd7謳n\x82־\xb3\x8f\xce\xe4\xdbϙ\xe2\xe0\xa4\x04\xf1:|\xc8A&\x85w\x15RN0e\xdd9\x0en\U0007b1c2\xb9\xba\xb1\x86,\xe1\xb1\f\x99\x8a\xd9\xffL\xe3\x01k嗫\xb5\xf2\xc0\x9dbi\xa2_a\xf9Zc\x9dl\x82\xba\x02{\x1dR\x95\xa9\xda\xc7Ş3]\xb6Z\x1d\xa9\x8e\xe05\xb7\x15\xc84\xf1Ǫ\xa17\xefa\x8f\x8bm|\xb6\xc9D\x0ej\x96\xfe\x9d\xe4\xb8'\x1eD\xb6u\xe8m\xbeq/h\xbc~\f\x99&\x14\x1a\xbf}\xfb{4\xba\x14\xf9nbB\x90\xc2\x11iQ\xf4\xfc!,Jo\x83\x1eA\nMD\x17D\xef#;\x17\xaf\a\xc1N\x1dΏ\x8c\xfe1h]\x9a\xf9\x8c\xf7\vF\xeb\xc4Z\x00\xf9MH\xeb\x17\xb2\x82\xb3\x9f\xe0t\x93}).\xe4\xf7\x97\xdd\x1c\x1dT Ys\x9f\x13\xb5J\x94\x94Y\xfdj~<\xb0\x1d\xfb<Ȕ\xb3x\xed\x93#\n;kRN\x8cTn7\xed\xb2\x9b#\xf1\xb6\x13\xaa\xd8\x00\xbdÂ\xc0Q\x1c*\xfd\xcbn\x99\xe3\xbeys]N\x88\x9f\x0f;\xb5̇\xe6\xecr\x96\xf9?'x1\xf8\x9c\xf7\x8b^\xff\xd3\xe7\xc1\xa9g~\xd1\x17\x98}&\x04\x02\xfc\xec\xac \xf0\xfe)\xafg)<\xf7\x88\x11\xa4j\u007f\xa8$\xb6\x03\x99\x92\xfeW\xa6\xae\xe8\xac\\F˽\x8a\x99u\xf8\x8d\xb5>B6\x9eۙ\xdba\x1f\x00i6\xf0\xe4)\xff\xc4?#N\xfa&\x03>!\x10\xf0\xf23~o0\x18\x9c\x9a\xf2\xfbff|\f\xd2\u007f\x01\x00\x00\xff\xff\t\x13eh/\b\x00\x00",
		hash:  "44aab25c24e40aa95ba19abe9cdfb21df8215049f123077a35e21f740ed91e37",
		mime:  "",
		mtime: time.Unix(1570537390, 0),
		size:  2095,
	},
	"PersistentEvent.md": {
		data:  "\x1f\x8b\b\x00\x00\x00\x00\x00\x02\xff\xd4VKs\x1aG\x17\xdd\xf3+\xfa+m>WY\b\x900\x90\x9d\xe3\xf2\xcaI\xcaUNV\xae,\x10\x8c\x1c\x97$P\fN\xcaU,\x86\x11\x83x\nH\xf4@\x80\x8c\xb0\x8d\x84\x1f0(\x960\xcc\xf0\xf81\xf4\xednV\xfa\v\xa9\x99F\x88 \xc5e/Ê\xea\xee{\xef\xb9\xe7\x9c\xdb=ss\xe8\xa1\xf0,\xf04\x10\x14|\xc1\xfb\xbf\t\xbe S\xceH~\xdbd\x9a\x9b\x9bCT\x8b`\xad\x85\xb5$\xd5\xd2&Ӄ\xe7\xcb\xc23\x9f\x10\x14\x02\xc88\x19@\x90\x92I\xf2\x03V\xa3\x88\xef\xe9[\xa3b\x94\xf6\xdf\xd0B\x84\r\xb2\xac\x92\x82?R\x90\x8d\xd0S\x8d\x9dG\x88\x9a\xd5כ\x9b\xa0\x1e\x93M\x19\xa2\xe7\x17\xdd\x14(\x1d\xaa\x9d\x91\xc4\x1b\xac\xaa\x90˓\xfd*\xb4\xdb\xecD\x02\xa5<\t\x81L\x9e\x14ϰV\x85L\x02\x1a\x05\xdcK\x8f\xf6\x1b\xa3Wy8\xdc\"\xa58\xb4\x1b\x90x;\x14\xa5\x1f\x1f\xdcGdG!\xa90n\xabD\xaa\xd1B\x84\xc4ER\x8asL#9M{\r\xac&\xb1\xd6\xd2\xcftdH\xedA\xa2\xcc6{:\x8e\xae\bY\x85\x94>\xf0\x15\xc8n_tS\xb8[\x80f\x94\xa7\xe2Ix84^\x92\xfd\x16(]\xd8R!\xd6\xe4\x99u\xa0\xf5<H5\xfa^\x19\xa3a\x83\xfc\x18Po@wjL\uecc6\x82\xd5\x1c\xc9dq\xbfH\v\x11\xf4\xf0\xee\xddG\x88\x94Ґ\xa8\x90\xd8\x1etE\xa2fYO\xd1k(\x9d1\xd6\xf30i\x9e\xb1A\x91UR\xe4\xb5H\xcf^\rE\x89\v4ڬ\xd1\xde_P\xaaqp\x10\x95W'\"\x81\xd2a\xa7\x15\x93iz\x97\x9f\x9fQ\x1c\xdd\xf5z\xe7\xfd>t\x1bA3z\xfd4֪\xb8\x9d\xbc\x9e7\x84f\xd7 \x9b\xa6'M\x14B\xf4T\x83\x97I\x14B#-\xcf\x1ao }\xa47o\xa8\x89B\x88\xc4E8}\xf9\x83{]\bl\xb8=B\x00\x85L!4?\xfbC\xc6\x1a\x9a\xdeA3\xc7\xf4\xc0\xe0\xaa0\xbf1\xe9g^\xd0\x1b\ny\x85\x8d5\xff\x8bu\xfd\xaf\xc5l#G\xed{\x0f\u007f\xbam\xb5X\xbe\xff\x16\xa22\xd4\xf3!\x1d\xf9|\xe0E (\xac\x87t&g\x87\x80\xcb\x05%\x95\x1c(\xd3\xc6\x1f\v\x92W\xa69\x1a\xdbt\xa7\f\xb9\x1e\xe4j\xa4T\x1e\x8aa\xbeȽ>\x14\xc3Е\xa0݆\\\x82Տi=\xceݏ\xd5\x13\xbaS\xd6G\xa2фޮn\u008c\x82\xb5\xeaH,\xb0\xc1ָ֧S\x90\xabX\x8d\xb2\x93\xd7\xd3E!\x93\xe3\xb1XM\x92MY\xf7j?\t\xc7\x12\xafHkIP3\x10.\xc0\x96\xaa\xdb2\xfa\x91\xd6\xf7Hl\x8f\x892)g\xb1Z\xdc\xf0{\xb1\xb6\xcd\x12\x12\x95:\xb8]g\xafޏ\xc4\xec\xf8\xf0v\x19\x8aG\x86Ů\x04\xe6BbM\xe6\xceŃC\xba{\x80\xdb\"n\xbf\x83f\x86\xec\xb7h!\xb2:C\xd3P\x94fyE\xe3p5\n\xcd\xe8\xd4}2;\x94|\x8c\n\x11\x88\xeb\x83u5\xa3r\x985\xdaD\xaaq\x86f\x93C3:\xa1h<\xa7\x97\x13J\xa4\x1a\xdbzGJ\xf1\x99!\xbdI\xfc\xd1A\x16b-rX\xc1Z\xcbd\xb2\x9a\x114\xe2\xec\xb5|C5nl^\xf3F\xa3\xe9\xda\x1aF0J\xd9\xcc\bk2\x94jV\xb3\x93\xc6c\xa4\xf4\xc1\x98\xad\xc4?\xb8\xe3\xd9\xf8\xad\xf1/\x00\xb9;\xc9^\x87|\xdc\xe5\xf7\x00G\b\x9d\x16k\xf4u\xc2.;4\xe0\xd3\x03\rz\xbb\x8f\xa1с\x83\xda\xf8\xaa\xd9>\x81X\v2͟\xff\xffK0\xb8\x11\xf8fa\xc1\xe3\xf7\x05\xfck\x82\xf9WϚ\xff\xb9\xd7\xec\xf1\xaf/\x04W\x05ۭK\xe4P\xaa\xc1\xa7c\xdc?\x01\xa5\xcbb\xef\xc8Q\x06\xb7\xeb\xba&\xe9]\xd8҆b\x96\xc4\xdf\xc2\xe9.\u007f2\x86b\ue89bb\x83\"\xc8\xd5\xe9uڨ\xd0ltT9\x1f\x1d\xf2KlьFb\x9c$ߎJ\";\x0e\xf3>h!2+\x89A\xcaE7E\xa5θ\xda^\x134\x95ׁc\t\x8a\xfd\x8bn\xc1d\xfa\xdf㫎\xd6\xddO}\xe3v\x9e\xae?1:z\xe6\xfe}\xc1\xe6\xb1z]\x8e\xa5e\xfb\xb2}i\xd1\xe1\xf6:m\x8b\xcb\xceE\xc1-\xd8\xef\xd8\x05\x8f\xcbn\xde\xf0=\xb9e2-\x99\xd1\xcd\x0fƌ{\xe6\x10)\x9e\xe9h\xfe\xab\x9cs\xf8\x9f\xe1\x9c\x1f\x9eٽ\x12\xc2\b\xe7T]\x93\xe3\xcbI\x84R\xed\xdew\x8f&\x1c\xf1w\x8eGM\x13\xaa\u007f\x1e\f\xf6\xaf\x13\xfa\x19F=k\x81\xaf \xd4ȯ\xb7n\xb0\xc6\x1b\xba\x91\x80ɛ?\t\xb9\xe8\xa6H\xfcO\xfd\t\xe7 \x8d\x16\xf8\xdde\xd4\xff\x02o\nv\xbb\xc5\xe5X\xb2\xaf\xac\xb8\xed6\xefʢ\xcb氹\x1d\x1e\x97\xc3\xear\xb8\x9d^\xe7؛\x13\x05\x99R\x9d.\xcfE\x99\xe6\xee\xeb\xa7\xc3aq\xacح.\x8f}\xe5\x8e\xd3v\xc7auY\\\x82e\xd1jw:\x9c\xcbv\xc7\xe5t\xfc\x1d\x00\x00\xff\xff\x18\n\x0e\xb18\n\x00\x00",
		hash:  "9e1b894cd8de723613d96dcf222679e753f3c27ce166b75f66e2ba7c1576b018",
		mime:  "",
		mtime: time.Unix(1570537390, 0),
		size:  2616,
	},
	"Prometheus.md": {
		data:  "",
		hash:  "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
		mime:  "",
		mtime: time.Unix(1570537390, 0),
		size:  0,
	},
	"TappController.md": {
		data:  "",
		hash:  "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
		mime:  "",
		mtime: time.Unix(1570537390, 0),
		size:  0,
	},
	"VolumeDecorator.md": {
		data:  "",
		hash:  "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
		mime:  "",
		mtime: time.Unix(1570537390, 0),
		size:  0,
	},
}

// NotFound is called when no asset is found.
// It defaults to http.NotFound but can be overwritten
var NotFound = http.NotFound

// ServeHTTP serves a request, attempting to reply with an embedded file.
func ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	f, ok := staticFiles[strings.TrimPrefix(req.URL.Path, "/")]
	if !ok {
		NotFound(rw, req)
		return
	}
	header := rw.Header()
	if f.hash != "" {
		if hash := req.Header.Get("If-None-Match"); hash == f.hash {
			rw.WriteHeader(http.StatusNotModified)
			return
		}
		header.Set("ETag", f.hash)
	}
	if !f.mtime.IsZero() {
		if t, err := time.Parse(http.TimeFormat, req.Header.Get("If-Modified-Since")); err == nil && f.mtime.Before(t.Add(1*time.Second)) {
			rw.WriteHeader(http.StatusNotModified)
			return
		}
		header.Set("Last-Modified", f.mtime.UTC().Format(http.TimeFormat))
	}
	header.Set("Content-Type", f.mime)

	// Check if the asset is compressed in the binary
	if f.size == 0 {
		header.Set("Content-Length", strconv.Itoa(len(f.data)))
		io.WriteString(rw, f.data)
	} else {
		if header.Get("Content-Encoding") == "" && strings.Contains(req.Header.Get("Accept-Encoding"), "gzip") {
			header.Set("Content-Encoding", "gzip")
			header.Set("Content-Length", strconv.Itoa(len(f.data)))
			io.WriteString(rw, f.data)
		} else {
			header.Set("Content-Length", strconv.Itoa(f.size))
			reader, _ := gzip.NewReader(strings.NewReader(f.data))
			io.Copy(rw, reader)
			reader.Close()
		}
	}
}

// Server is simply ServeHTTP but wrapped in http.HandlerFunc so it can be passed into net/http functions directly.
var Server http.Handler = http.HandlerFunc(ServeHTTP)

// Open allows you to read an embedded file directly. It will return a decompressing Reader if the file is embedded in compressed format.
// You should close the Reader after you're done with it.
func Open(name string) (io.ReadCloser, error) {
	f, ok := staticFiles[name]
	if !ok {
		return nil, fmt.Errorf("Asset %s not found", name)
	}

	if f.size == 0 {
		return ioutil.NopCloser(strings.NewReader(f.data)), nil
	}
	return gzip.NewReader(strings.NewReader(f.data))
}

// ModTime returns the modification time of the original file.
// Useful for caching purposes
// Returns zero time if the file is not in the bundle
func ModTime(file string) (t time.Time) {
	if f, ok := staticFiles[file]; ok {
		t = f.mtime
	}
	return
}

// Hash returns the hex-encoded SHA256 hash of the original file
// Used for the Etag, and useful for caching
// Returns an empty string if the file is not in the bundle
func Hash(file string) (s string) {
	if f, ok := staticFiles[file]; ok {
		s = f.hash
	}
	return
}
