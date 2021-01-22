package zones

// Reference:
// Japanese: https://jp.cybozu.help/general/ja/admin/list_systemadmin/list_system_time/timezone.html
// Chinese: https://jp.cybozu.help/general/zh/admin/list_systemadmin/list_system_time/timezone.html
// English: https://jp.cybozu.help/general/en/admin/list_systemadmin/list_system_time/timezone.html

const (
	UTC   = "UTC"   // time.UTC
	Local = "Local" // time.Local

	Etc_GMT_P12                    = "Etc/GMT+12"                     // (UTC-12:00) International Date Line West.							国际日期变更线西.										国際日付変更線 西側.
	Etc_GMT_P11                    = "Etc/GMT+11"                     // (UTC-11:00) Coordinated Universal Time-11.							协调世界时-11.										協定世界時-11.
	Pacific_Honolulu               = "Pacific/Honolulu"               // (UTC-10:00) Hawaii.												夏威夷.												ハワイ.
	America_Anchorage              = "America/Anchorage"              // (UTC-09:00) Alaska.												安克雷奇.											アラスカ.
	America_Santa_Isabel           = "America/Santa_Isabel"           // (UTC-08:00) Baja California.										下加利福尼亚州.										バハカリフォルニア.
	America_Los_Angeles            = "America/Los_Angeles"            // (UTC-08:00) Pacific Time (US and Canada).							太平洋时间（美国和加拿大）.							太平洋標準時(米国およびカナダ).
	America_Chihuahua              = "America/Chihuahua"              // (UTC-07:00) Chihuahua, La Paz, Mazatlan.							奇瓦瓦,拉巴斯,马萨特兰.								チワワ、ラパス、マサトラン.
	America_Phoenix                = "America/Phoenix"                // (UTC-07:00) Arizona.												亚利桑那.											アリゾナ.
	America_Denver                 = "America/Denver"                 // (UTC-07:00) Mountain Time (US and Canada).							山地时间（美国和加拿大）.								山地標準時(米国およびカナダ).
	America_Guatemala              = "America/Guatemala"              // (UTC-06:00) Central America.										中美洲.												中央アメリカ.
	America_Chicago                = "America/Chicago"                // (UTC-06:00) Central Time (US and Canada).							中部时间（美国和加拿大）.								中部標準時(米国およびカナダ).
	America_Regina                 = "America/Regina"                 // (UTC-06:00) Saskatchewan.											萨斯喀彻温.											サスカチュワン.
	America_Mexico_City            = "America/Mexico_City"            // (UTC-06:00) Guadalajara, Mexico City, Monterey.					瓜达拉哈拉,墨西哥城,蒙特雷.							グアダラハラ、メキシコシティ、モンテレー.
	America_Bogota                 = "America/Bogota"                 // (UTC-05:00) Bogota, Lima, Quito.									波哥大,利马,基多.										ボゴタ、リマ、キト.
	America_Indiana_Indianapolis   = "America/Indiana/Indianapolis"   // (UTC-05:00) Indiana (East).										印地安那州（东部）.									インディアナ東部.
	America_New_York               = "America/New_York"               // (UTC-05:00) Eastern Time (US and Canada).							东部时间（美国和加拿大）.								東部標準時(米国およびカナダ).
	America_Caracas                = "America/Caracas"                // (UTC-04:30) Caracas.												加拉加斯.											カラカス.
	America_Halifax                = "America/Halifax"                // (UTC-04:00) Atlantic Time (Canada).								大西洋时间（加拿大）.									大西洋標準時(カナダ).
	America_Asuncion               = "America/Asuncion"               // (UTC-04:00) Asuncion.												亚松森.												アスンシオン.
	America_La_Paz                 = "America/La_Paz"                 // (UTC-04:00) Georgetown, La Paz, Manaus, San Juan.					乔治敦,拉巴斯,马瑙斯,圣胡安.							ジョージタウン、ラパス、マナウス、サンフアン.
	America_Cuiaba                 = "America/Cuiaba"                 // (UTC-04:00) Cuiaba.												库亚巴.												クイアバ.
	America_Santiago               = "America/Santiago"               // (UTC-04:00) Santiago.												圣地亚哥.											サンチアゴ.
	America_St_Johns               = "America/St_Johns"               // (UTC-03:30) Newfoundland.											纽芬兰.												ニューファンドランド.
	America_Sao_Paulo              = "America/Sao_Paulo"              // (UTC-03:00) Brasilia.												巴西利亚.											ブラジリア.
	America_Godthab                = "America/Godthab"                // (UTC-03:00) Greenland.												格陵兰.												グリーンランド.
	America_Cayenne                = "America/Cayenne"                // (UTC-03:00) Cayenne, Fortaleza.									卡宴,福塔雷萨.										カイエンヌ、フォルタレザ.
	America_Argentina_Buenos_Aires = "America/Argentina/Buenos_Aires" // (UTC-03:00) Buenos Aires.											布宜诺斯艾利斯.										ブエノスアイレス.
	America_Montevideo             = "America/Montevideo"             // (UTC-03:00) Montevideo.											蒙得维的亚.											モンテビデオ.
	Etc_GMT_P2                     = "Etc/GMT+2"                      // (UTC-02:00) Coordinated Universal Time-2.							协调世界时-2.										協定世界時-2.
	Atlantic_Cape_Verde            = "Atlantic/Cape_Verde"            // (UTC-01:00) Cape Verde.											佛得角群岛.											カーボベルデ諸島.
	Atlantic_Azores                = "Atlantic/Azores"                // (UTC-01:00) Azores.												亚速尔群岛.											アゾレス.
	Africa_Casablanca              = "Africa/Casablanca"              // (UTC+00:00) Casablanca.											卡萨布兰卡.											カサブランカ.
	Atlantic_Reykjavik             = "Atlantic/Reykjavik"             // (UTC+00:00) Monrovia, Reykjavik.									蒙罗维亚,雷克雅未克.									モンロビア、レイキャビク.
	Europe_London                  = "Europe/London"                  // (UTC+00:00) Dublin, Edinburgh, Lisbon, London.						都柏林,爱丁堡,里斯本,伦敦.								ダブリン、エジンバラ、リスボン、ロンドン.
	Etc_GMT                        = "Etc/GMT"                        // (UTC+00:00) Coordinated Universal Time.							协调世界时.											協定世界時.
	Europe_Berlin                  = "Europe/Berlin"                  // (UTC+01:00) Amsterdam, Berlin, Bern, Rome, Stockholm, Vienna.		阿姆斯特丹,柏林,伯尔尼,罗马,斯德哥尔摩,维也纳.			アムステルダム、ベルリン、ベルン、ローマ、ストックホルム、ウィーン.
	Europe_Paris                   = "Europe/Paris"                   // (UTC+01:00) Brussels, Copenhagen, Madrid, Paris.					布鲁塞尔,哥本哈根,马德里,巴黎.							ブリュッセル、コペンハーゲン、マドリード、パリ.
	Africa_Lagos                   = "Africa/Lagos"                   // (UTC+01:00) West Central Africa.									中非西部.											西中央アフリカ.
	Europe_Budapest                = "Europe/Budapest"                // (UTC+01:00) Belgrade, Bratislava, Budapest, Ljubljana, Prague.		贝尔格莱德,布拉迪斯拉发,布达佩斯,卢布尔雅那,布拉格.		ベオグラード、ブラチスラバ、ブダペスト、リュブリャナ、プラハ.
	Europe_Warsaw                  = "Europe/Warsaw"                  // (UTC+01:00) Sarajevo, Skopje, Warsaw, Zagreb.						萨拉热窝,斯科普里,华沙,萨格勒布.						サラエボ、スコピエ、ワルシャワ、ザグレブ.
	Africa_Windhoek                = "Africa/Windhoek"                // (UTC+01:00) Windhoek.												温得和克.											ウィントフック.
	Europe_Istanbul                = "Europe/Istanbul"                // (UTC+02:00) Athens, Bucharest, Istanbul.							雅典,布加勒斯特,伊斯坦布尔.							アテネ、ブカレスト、イスタンブール.
	Europe_Kiev                    = "Europe/Kiev"                    // (UTC+02:00) Helsinki, Kiev, Riga, Sofia, Tallinn, Vilnius.			赫尔辛基,基辅,里加,索非亚,塔林,维尔纽斯.				ヘルシンキ、キエフ、リガ、ソフィア、タリン、ビリニュス.
	Africa_Cairo                   = "Africa/Cairo"                   // (UTC+02:00) Cairo.													开罗.												カイロ.
	Asia_Damascus                  = "Asia/Damascus"                  // (UTC+02:00) Damascus.												大马士革.											ダマスカス.
	Asia_Amman                     = "Asia/Amman"                     // (UTC+02:00) Amman.													安曼.												アンマン.
	Africa_Johannesburg            = "Africa/Johannesburg"            // (UTC+02:00) Harare, Pretoria.										哈拉雷,比勒陀利亚.									ハラーレ、プレトリア.
	Asia_Jerusalem                 = "Asia/Jerusalem"                 // (UTC+02:00) Jerusalem.												耶路撒冷.											エルサレム.
	Asia_Beirut                    = "Asia/Beirut"                    // (UTC+02:00) Beirut.												贝鲁特.												ベイルート.
	Asia_Baghdad                   = "Asia/Baghdad"                   // (UTC+03:00) Baghdad.												巴格达.												バグダッド.
	Europe_Minsk                   = "Europe/Minsk"                   // (UTC+03:00) Minsk.													明斯克.												ミンスク.
	Asia_Riyadh                    = "Asia/Riyadh"                    // (UTC+03:00) Kuwait, Riyadh.										利雅得.												クエート、リヤド.
	Africa_Nairobi                 = "Africa/Nairobi"                 // (UTC+03:00) Nairobi.												内罗毕.												ナイロビ.
	Asia_Tehran                    = "Asia/Tehran"                    // (UTC+03:30) Tehran.												德黑兰.												テヘラン.
	Europe_Moscow                  = "Europe/Moscow"                  // (UTC+04:00) Moscow, St. Petersburg, Volgograd.						莫斯科,圣彼得堡,伏尔加格勒.							モスクワ、サンクトペテルブルグ、ボルゴグラード.
	Asia_Tbilisi                   = "Asia/Tbilisi"                   // (UTC+04:00) Tbilisi.												第比利斯.											トビリシ.
	Asia_Yerevan                   = "Asia/Yerevan"                   // (UTC+04:00) Yerevan.												埃里温.												エレバン.
	Asia_Dubai                     = "Asia/Dubai"                     // (UTC+04:00) Abu Dhabi, Muscat.										阿布扎比,马斯喀特.									アブダビ、マスカット.
	Asia_Baku                      = "Asia/Baku"                      // (UTC+04:00) Baku.													巴库.												バクー.
	Indian_Mauritius               = "Indian/Mauritius"               // (UTC+04:00) Port Louis.											路易港.												ポートルイス.
	Asia_Kabul                     = "Asia/Kabul"                     // (UTC+04:30) Kabul.													喀布尔.												カブール.
	Asia_Tashkent                  = "Asia/Tashkent"                  // (UTC+05:00) Tashkent.												塔什干.												タシケント.
	Asia_Karachi                   = "Asia/Karachi"                   // (UTC+05:00) Islamabad, Karachi.									伊斯兰堡,卡拉奇.										イスラマバード、カラチ.
	Asia_Colombo                   = "Asia/Colombo"                   // (UTC+05:30) Sri Jayewardenepura Kotte.								斯里加亚渥登普拉.										スリジャヤワルダナプラコッテ.
	Asia_Kolkata                   = "Asia/Kolkata"                   // (UTC+05:30) Chennai, Kolkata, Mumbai, New Delhi.					钦奈,加尔各答,孟买,新德里.								チェンナイ、コルカタ、ムンバイ、ニューデリー.
	Asia_Kathmandu                 = "Asia/Kathmandu"                 // (UTC+05:45) Kathmandu.												加德满都.											カトマンズ.
	Asia_Almaty                    = "Asia/Almaty"                    // (UTC+06:00) Astana.												阿斯塔纳.											アスタナ.
	Asia_Dhaka                     = "Asia/Dhaka"                     // (UTC+06:00) Dhaka.													达卡.												ダッカ.
	Asia_Yekaterinburg             = "Asia/Yekaterinburg"             // (UTC+06:00) Yekaterinburg.											叶卡捷琳堡.											エカテリンブルグ.
	Asia_Yangon                    = "Asia/Yangon"                    // (UTC+06:30) Yangon.												仰光.												ヤンゴン.
	Asia_Bangkok                   = "Asia/Bangkok"                   // (UTC+07:00) Bangkok, Hanoi, Jakarta.								曼谷,河内,雅加达.										バンコク、ハノイ、ジャカルタ.
	Asia_Novosibirsk               = "Asia/Novosibirsk"               // (UTC+07:00) Novosibirsk.											新西伯利亚.											ノヴォシビルスク.
	Asia_Krasnoyarsk               = "Asia/Krasnoyarsk"               // (UTC+08:00) Krasnoyarsk.											克拉斯诺亚尔斯克.										クラスノヤルスク.
	Asia_Ulaanbaatar               = "Asia/Ulaanbaatar"               // (UTC+08:00) Ulaanbaatar.											乌兰巴托.											ウランバートル.
	Asia_Shanghai                  = "Asia/Shanghai"                  // (UTC+08:00) Beijing, Chongqing, Hong Kong, Urumqi.					北京,重庆,香港,乌鲁木齐.								北京、重慶、香港、ウルムチ.
	Australia_Perth                = "Australia/Perth"                // (UTC+08:00) Perth.													佩思.												パース.
	Asia_Singapore                 = "Asia/Singapore"                 // (UTC+08:00) Kuala Lumpur, Singapore.								吉隆坡,新加坡.										クアラルンプール、シンガポール.
	Asia_Taipei                    = "Asia/Taipei"                    // (UTC+08:00) Taipei.												台北.												台北.
	Asia_Irkutsk                   = "Asia/Irkutsk"                   // (UTC+09:00) Irkutsk.												伊尔库茨克.											イルクーツク.
	Asia_Seoul                     = "Asia/Seoul"                     // (UTC+09:00) Seoul.													首尔.												ソウル.
	Asia_Tokyo                     = "Asia/Tokyo"                     // (UTC+09:00) Osaka, Sapporo, Tokyo.									大阪,札幌,东京.										大阪、札幌、東京.
	Australia_Darwin               = "Australia/Darwin"               // (UTC+09:30) Darwin.												达尔文.												ダーウィン.
	Australia_Adelaide             = "Australia/Adelaide"             // (UTC+09:30) Adelaide.												阿德莱德.											アデレード.
	Australia_Hobart               = "Australia/Hobart"               // (UTC+10:00) Hobart.												霍巴特.												ホバート.
	Asia_Yakutsk                   = "Asia/Yakutsk"                   // (UTC+10:00) Yakutsk.												雅库茨克.											ヤクーツク.
	Australia_Brisbane             = "Australia/Brisbane"             // (UTC+10:00) Brisbane.												布里斯班.											ブリスベン.
	Pacific_Port_Moresby           = "Pacific/Port_Moresby"           // (UTC+10:00) Guam, Port Moresby.									关岛,莫尔兹比港.										グアム、ポートモレスビー.
	Australia_Sydney               = "Australia/Sydney"               // (UTC+10:00) Canberra, Melbourne, Sydney.							堪培拉,墨尔本,悉尼.									キャンベラ、メルボルン、シドニー.
	Asia_Vladivostok               = "Asia/Vladivostok"               // (UTC+11:00) Vladivostok.											符拉迪沃斯托克.										ウラジオストク.
	Pacific_Guadalcanal            = "Pacific/Guadalcanal"            // (UTC+11:00) Solomon Islands, New Caledonia.						所罗门群岛,新喀里多尼亚.								ソロモン諸島、ニューカレドニア.
	Etc_GMT_M12                    = "Etc/GMT-12"                     // (UTC+12:00) Coordinated Universal Time+12.							协调世界时+12.										協定世界時+12.
	Pacific_Fiji                   = "Pacific/Fiji"                   // (UTC+12:00) Fiji, Marshall Islands.								斐济,马绍尔群岛.										フィジー、マーシャル諸島.
	Asia_Magadan                   = "Asia/Magadan"                   // (UTC+12:00) Magadan.												马加丹.												マガダン.
	Pacific_Auckland               = "Pacific/Auckland"               // (UTC+12:00) Auckland, Wellington.									奥克兰,惠灵顿.										オークランド、ウェリントン.
	Pacific_Tongatapu              = "Pacific/Tongatapu"              // (UTC+13:00) Nuku'alofa.											努库阿洛法.											ヌクアロファ.
	Pacific_Apia                   = "Pacific/Apia"                   // (UTC+13:00) Samoa.													萨摩亚群岛.											サモア.
)
