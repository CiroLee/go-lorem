package data

var PROTOCOL []string = []string{
	"http",
	"https",
	"ftp",
	"gopher",
	"mailto",
	"mid",
	"cid",
	"news",
	"nntp",
	"prospero",
	"telnet",
	"rlogin",
	"tn3270",
	"wais",
}

var DOMAINS map[string]([]string) = map[string]([]string){
	"iTLD": []string{
		".com",
		".edu",
		".gov",
		".int",
		".mil",
		".net",
		".org",
		".biz",
		".info",
		".pro",
		".name",
		".museum",
		".coop",
		".aero",
		".idv",
	},
	"ccTLD": []string{
		".al",
		".dz",
		".af",
		".ar",
		".ae",
		".aw",
		".om",
		".az",
		".eg",
		".et",
		".ie",
		".ee",
		".ad",
		".ao",
		".ai",
		".ag",
		".at",
		".au",
		".mo",
		".bb",
		".pg",
		".bs",
		".pk",
		".py",
		".ps",
		".bh",
		".pa",
		".br",
		".by",
		".bm",
		".bg",
		".mp",
		".bj",
		".be",
		".is",
		".pr",
		".ba",
		".pl",
		".bo",
		".bz",
		".bw",
		".bt",
		".bf",
		".bi",
		".bv",
		".dk",
		".de",
		".tl",
		".tg",
		".dm",
		".do",
		".kp",
		".gq",
		".ru",
		".ec",
		".er",
		".fr",
		".fo",
		".pf",
		".gf",
		".tf",
		".va",
		".ph",
		".fj",
		".fi",
		".cv",
		".fk",
		".gm",
		".cg",
		".cd",
		".co",
		".cr",
		".gg",
		".gd",
		".gl",
		".ge",
		".cu",
		".gp",
		".gu",
		".gy",
		".kz",
		".ht",
		".kr",
		".nl",
		".an",
		".hm",
		".hn",
		".ki",
		".dj",
		".kg",
		".gn",
		".gw",
		".ca",
		".gh",
		".ga",
		".kh",
		".cz",
		".zw",
		".cm",
		".qa",
		".ky",
		".km",
		".ci",
		".kw",
		".cc",
		".hr",
		".ke",
		".ck",
		".lv",
		".ls",
		".la",
		".lb",
		".lt",
		".lr",
		".ly",
		".li",
		".re",
		".lu",
		".rw",
		".ro",
		".mg",
		".im",
		".mv",
		".mt",
		".mw",
		".my",
		".ml",
		".mk",
		".mh",
		".mq",
		".yt",
		".mu",
		".mr",
		".us",
		".um",
		".as",
		".vi",
		".mn",
		".ms",
		".bd",
		".pe",
		".fm",
		".mm",
		".md",
		".ma",
		".mc",
		".mz",
		".mx",
		".nr",
		".np",
		".ni",
		".ne",
		".ng",
		".nu",
		".no",
		".nf",
		".na",
		".za",
		".aq",
		".gs",
		".eu",
		".pw",
		".pn",
		".pt",
		".jp",
		".se",
		".ch",
		".sv",
		".ws",
		".yu",
		".sl",
		".sn",
		".cy",
		".sc",
		".sa",
		".cx",
		".st",
		".sh",
		".kn",
		".lc",
		".sm",
		".pm",
		".vc",
		".lk",
		".sk",
		".si",
		".sj",
		".sz",
		".sd",
		".sr",
		".sb",
		".so",
		".tj",
		".tw",
		".th",
		".tz",
		".to",
		".tc",
		".tt",
		".tn",
		".tv",
		".tr",
		".tm",
		".tk",
		".wf",
		".vu",
		".gt",
		".ve",
		".bn",
		".ug",
		".ua",
		".uy",
		".uz",
		".es",
		".eh",
		".gr",
		".hk",
		".sg",
		".nc",
		".nz",
		".hu",
		".sy",
		".jm",
		".am",
		".ac",
		".ye",
		".iq",
		".ir",
		".il",
		".it",
		".in",
		".id",
		".uk",
		".vg",
		".io",
		".jo",
		".vn",
		".zm",
		".je",
		".td",
		".gi",
		".cl",
		".cf",
		".cn",
	},
}

var EMAIL_SUFFIX []string = []string{
	"@gmail.com",
	"@yahoo.com",
	"@msn.com",
	"@hotmail.com",
	"@aol.com",
	"@ask.com",
	"@live.com",
	"@qq.com",
	"@vip.qq.com",
	"@foxmail.com",
	"@0355.net",
	"@163.net",
	"@263.net",
	"@yeah.net",
	"@googlemail.com",
	"@mail.com",
	"@aim.com",
	"@walla.com",
	"@inbox.com",
	"@126.com",
	"@163.com",
	"@139.com",
	"@189.cn",
	"@wo.cn",
	"@sina.com",
	"@21cn.com",
	"@sohu.com",
	"@yahoo.com.cn",
	"@tom.com",
	"@etang.com",
	"@eyou.com",
	"@56.com",
	"@x.cn",
	"@chinaren.com",
	"@sogou.com",
	"@citiz.com",
	"@hongkong.com",
	"@ctimail.com",
	"@hknet.com",
	"@netvigator.com",
	"@mail.hk.com",
	"@swe.com.hk",
	"@itccolp.com.hk",
	"@biznetvigator.com",
	"@seed.net.tw",
	"@topmarkeplg.com.tw",
}

var MESH_NUM []int = []int{
	139, 138, 137, 136, 135, 134, 159, 158, 157, 150, 151, 152, 188, 187, 182, 183, 184, 178, 130, 131, 132, 156, 155,
	186, 185, 176, 133, 153, 189, 180, 181, 177,
}

var TEL_AREA_NUM []string = []string{
	"010",
	"020",
	"021",
	"022",
	"023",
	"024",
	"025",
	"027",
	"028",
	"029",
	"0310",
	"0311",
	"0312",
	"0313",
	"0314",
	"0315",
	"0316",
	"0317",
	"0318",
	"0319",
	"0335",
	"0349",
	"0350",
	"0351",
	"0352",
	"0353",
	"0354",
	"0355",
	"0356",
	"0357",
	"0358",
	"0359",
	"0370",
	"0371",
	"0372",
	"0373",
	"0374",
	"0375",
	"0376",
	"0377",
	"0378",
	"0379",
	"0391",
	"0392",
	"0393",
	"0394",
	"0395",
	"0396",
	"0398",
	"0410",
	"0411",
	"0412",
	"0413",
	"0414",
	"0415",
	"0416",
	"0417",
	"0418",
	"0419",
	"0421",
	"0427",
	"0429",
	"0431",
	"0432",
	"0434",
	"0435",
	"0436",
	"0437",
	"0438",
	"0439",
	"0451",
	"0452",
	"0453",
	"0454",
	"0455",
	"0456",
	"0457",
	"0458",
	"0459",
	"0464",
	"0467",
	"0468",
	"0469",
	"0470",
	"0471",
	"0472",
	"0473",
	"0474",
	"0475",
	"0476",
	"0477",
	"0478",
	"0479",
	"0482",
	"0483",
	"0510",
	"0511",
	"0512",
	"0513",
	"0514",
	"0515",
	"0516",
	"0517",
	"0518",
	"0519",
	"0523",
	"0527",
	"0530",
	"0531",
	"0532",
	"0533",
	"0534",
	"0535",
	"0536",
	"0537",
	"0538",
	"0539",
	"0543",
	"0546",
	"0550",
	"0551",
	"0552",
	"0553",
	"0554",
	"0555",
	"0556",
	"0557",
	"0558",
	"0559",
	"0561",
	"0562",
	"0563",
	"0564",
	"0566",
	"0570",
	"0571",
	"0572",
	"0573",
	"0574",
	"0575",
	"0576",
	"0577",
	"0578",
	"0579",
	"0580",
	"0591",
	"0592",
	"0593",
	"0594",
	"0595",
	"0596",
	"0597",
	"0598",
	"0599",
	"0631",
	"0632",
	"0633",
	"0635",
	"0660",
	"0662",
	"0663",
	"0668",
	"0691",
	"0692",
	"0701",
	"0710",
	"0711",
	"0712",
	"0713",
	"0714",
	"0715",
	"0716",
	"0717",
	"0718",
	"0719",
	"0722",
	"0724",
	"0728",
	"0730",
	"0731",
	"0732",
	"0733",
	"0734",
	"0735",
	"0736",
	"0737",
	"0738",
	"0739",
	"0743",
	"0744",
	"0745",
	"0746",
	"0750",
	"0751",
	"0752",
	"0753",
	"0754",
	"0755",
	"0756",
	"0757",
	"0758",
	"0759",
	"0760",
	"0762",
	"0763",
	"0766",
	"0768",
	"0769",
	"0770",
	"0771",
	"0772",
	"0773",
	"0774",
	"0775",
	"0776",
	"0777",
	"0778",
	"0779",
	"0790",
	"0791",
	"0792",
	"0793",
	"0794",
	"0795",
	"0796",
	"0797",
	"0798",
	"0799",
	"0801",
	"0802",
	"0803",
	"0804",
	"0805",
	"0806",
	"0807",
	"0809",
	"0812",
	"0813",
	"0816",
	"0817",
	"0818",
	"0825",
	"0826",
	"0827",
	"0830",
	"0831",
	"0832",
	"0833",
	"0834",
	"0835",
	"0836",
	"0837",
	"0838",
	"0839",
	"0851",
	"0852",
	"0853",
	"0854",
	"0855",
	"0856",
	"0857",
	"0858",
	"0859",
	"0870",
	"0871",
	"0872",
	"0873",
	"0874",
	"0875",
	"0876",
	"0877",
	"0878",
	"0879",
	"0883",
	"0886",
	"0887",
	"0888",
	"0891",
	"0892",
	"0893",
	"0894",
	"0895",
	"0896",
	"0897",
	"0898",
	"0899",
	"0901",
	"0902",
	"0903",
	"0906",
	"0908",
	"0909",
	"0910",
	"0911",
	"0912",
	"0913",
	"0914",
	"0915",
	"0916",
	"0917",
	"0919",
	"0930",
	"0931",
	"0932",
	"0933",
	"0934",
	"0935",
	"0936",
	"0937",
	"0938",
	"0941",
	"0943",
	"0951",
	"0952",
	"0953",
	"0954",
	"0970",
	"0971",
	"0972",
	"0973",
	"0974",
	"0975",
	"0976",
	"0977",
	"0990",
	"0991",
	"0992",
	"0993",
	"0994",
	"0995",
	"0996",
	"0997",
	"0998",
	"0999",
	"1391",
	"1433",
	"1558",
	"1719",
	"1728",
	"1755",
	"1771",
	"1772",
	"1774",
	"1832",
	"1833",
	"1852",
	"1853",
	"1886",
	"1892",
	"1893",
	"1894",
	"1896",
	"1897",
	"1898",
	"1899",
	"1903",
	"1906",
	"1909",
	"1935",
	"1937",
	"1953",
	"1994",
	"1996",
	"1997",
	"1998",
	"1999",
	"2728",
	"2802",
	"2898",
	"2935",
}
