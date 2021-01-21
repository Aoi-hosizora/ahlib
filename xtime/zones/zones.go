package zones

// reference: https://jp.cybozu.help/general/ja/admin/list_systemadmin/list_system_time/timezone.html

const (
	Etc_GMT_P12                    = "Etc/GMT+12"                     // (UTC-12:00) 国際日付変更線 西側
	Etc_GMT_P11                    = "Etc/GMT+11"                     // (UTC-11:00) 協定世界時-11
	Pacific_Honolulu               = "Pacific/Honolulu"               // (UTC-10:00) ハワイ
	America_Anchorage              = "America/Anchorage"              // (UTC-09:00) アラスカ
	America_Santa_Isabel           = "America/Santa_Isabel"           // (UTC-08:00) バハカリフォルニア
	America_Los_Angeles            = "America/Los_Angeles"            // (UTC-08:00) 太平洋標準時(米国およびカナダ)
	America_Chihuahua              = "America/Chihuahua"              // (UTC-07:00) チワワ、ラパス、マサトラン
	America_Phoenix                = "America/Phoenix"                // (UTC-07:00) アリゾナ
	America_Denver                 = "America/Denver"                 // (UTC-07:00) 山地標準時(米国およびカナダ)
	America_Guatemala              = "America/Guatemala"              // (UTC-06:00) 中央アメリカ
	America_Chicago                = "America/Chicago"                // (UTC-06:00) 中部標準時(米国およびカナダ)
	America_Regina                 = "America/Regina"                 // (UTC-06:00) サスカチュワン
	America_Mexico_City            = "America/Mexico_City"            // (UTC-06:00) グアダラハラ、メキシコシティ、モンテレー
	America_Bogota                 = "America/Bogota"                 // (UTC-05:00) ボゴタ、リマ、キト
	America_Indiana_Indianapolis   = "America/Indiana/Indianapolis"   // (UTC-05:00) インディアナ東部
	America_New_York               = "America/New_York"               // (UTC-05:00) 東部標準時(米国およびカナダ)
	America_Caracas                = "America/Caracas"                // (UTC-04:30) カラカス
	America_Halifax                = "America/Halifax"                // (UTC-04:00) 大西洋標準時(カナダ)
	America_Asuncion               = "America/Asuncion"               // (UTC-04:00) アスンシオン
	America_La_Paz                 = "America/La_Paz"                 // (UTC-04:00) ジョージタウン、ラパス、マナウス、サンフアン
	America_Cuiaba                 = "America/Cuiaba"                 // (UTC-04:00) クイアバ
	America_Santiago               = "America/Santiago"               // (UTC-04:00) サンチアゴ
	America_St_Johns               = "America/St_Johns"               // (UTC-03:30) ニューファンドランド
	America_Sao_Paulo              = "America/Sao_Paulo"              // (UTC-03:00) ブラジリア
	America_Godthab                = "America/Godthab"                // (UTC-03:00) グリーンランド
	America_Cayenne                = "America/Cayenne"                // (UTC-03:00) カイエンヌ、フォルタレザ
	America_Argentina_Buenos_Aires = "America/Argentina/Buenos_Aires" // (UTC-03:00) ブエノスアイレス
	America_Montevideo             = "America/Montevideo"             // (UTC-03:00) モンテビデオ
	Etc_GMT_P2                     = "Etc/GMT+2"                      // (UTC-02:00) 協定世界時-2
	Atlantic_Cape_Verde            = "Atlantic/Cape_Verde"            // (UTC-01:00) カーボベルデ諸島
	Atlantic_Azores                = "Atlantic/Azores"                // (UTC-01:00) アゾレス
	Africa_Casablanca              = "Africa/Casablanca"              // (UTC+00:00) カサブランカ
	Atlantic_Reykjavik             = "Atlantic/Reykjavik"             // (UTC+00:00) モンロビア、レイキャビク
	Europe_London                  = "Europe/London"                  // (UTC+00:00) ダブリン、エジンバラ、リスボン、ロンドン
	Etc_GMT                        = "Etc/GMT"                        // (UTC+00:00) 協定世界時
	Europe_Berlin                  = "Europe/Berlin"                  // (UTC+01:00) アムステルダム、ベルリン、ベルン、ローマ、ストックホルム、ウィーン
	Europe_Paris                   = "Europe/Paris"                   // (UTC+01:00) ブリュッセル、コペンハーゲン、マドリード、パリ
	Africa_Lagos                   = "Africa/Lagos"                   // (UTC+01:00) 西中央アフリカ
	Europe_Budapest                = "Europe/Budapest"                // (UTC+01:00) ベオグラード、ブラチスラバ、ブダペスト、リュブリャナ、プラハ
	Europe_Warsaw                  = "Europe/Warsaw"                  // (UTC+01:00) サラエボ、スコピエ、ワルシャワ、ザグレブ
	Africa_Windhoek                = "Africa/Windhoek"                // (UTC+01:00) ウィントフック
	Europe_Istanbul                = "Europe/Istanbul"                // (UTC+02:00) アテネ、ブカレスト、イスタンブール
	Europe_Kiev                    = "Europe/Kiev"                    // (UTC+02:00) ヘルシンキ、キエフ、リガ、ソフィア、タリン、ビリニュス
	Africa_Cairo                   = "Africa/Cairo"                   // (UTC+02:00) カイロ
	Asia_Damascus                  = "Asia/Damascus"                  // (UTC+02:00) ダマスカス
	Asia_Amman                     = "Asia/Amman"                     // (UTC+02:00) アンマン
	Africa_Johannesburg            = "Africa/Johannesburg"            // (UTC+02:00) ハラーレ、プレトリア
	Asia_Jerusalem                 = "Asia/Jerusalem"                 // (UTC+02:00) エルサレム
	Asia_Beirut                    = "Asia/Beirut"                    // (UTC+02:00) ベイルート
	Asia_Baghdad                   = "Asia/Baghdad"                   // (UTC+03:00) バグダッド
	Europe_Minsk                   = "Europe/Minsk"                   // (UTC+03:00) ミンスク
	Asia_Riyadh                    = "Asia/Riyadh"                    // (UTC+03:00) クエート、リヤド
	Africa_Nairobi                 = "Africa/Nairobi"                 // (UTC+03:00) ナイロビ
	Asia_Tehran                    = "Asia/Tehran"                    // (UTC+03:30) テヘラン
	Europe_Moscow                  = "Europe/Moscow"                  // (UTC+04:00) モスクワ、サンクトペテルブルグ、ボルゴグラード
	Asia_Tbilisi                   = "Asia/Tbilisi"                   // (UTC+04:00) トビリシ
	Asia_Yerevan                   = "Asia/Yerevan"                   // (UTC+04:00) エレバン
	Asia_Dubai                     = "Asia/Dubai"                     // (UTC+04:00) アブダビ、マスカット
	Asia_Baku                      = "Asia/Baku"                      // (UTC+04:00) バクー
	Indian_Mauritius               = "Indian/Mauritius"               // (UTC+04:00) ポートルイス
	Asia_Kabul                     = "Asia/Kabul"                     // (UTC+04:30) カブール
	Asia_Tashkent                  = "Asia/Tashkent"                  // (UTC+05:00) タシケント
	Asia_Karachi                   = "Asia/Karachi"                   // (UTC+05:00) イスラマバード、カラチ
	Asia_Colombo                   = "Asia/Colombo"                   // (UTC+05:30) スリジャヤワルダナプラコッテ
	Asia_Kolkata                   = "Asia/Kolkata"                   // (UTC+05:30) チェンナイ、コルカタ、ムンバイ、ニューデリー
	Asia_Kathmandu                 = "Asia/Kathmandu"                 // (UTC+05:45) カトマンズ
	Asia_Almaty                    = "Asia/Almaty"                    // (UTC+06:00) アスタナ
	Asia_Dhaka                     = "Asia/Dhaka"                     // (UTC+06:00) ダッカ
	Asia_Yekaterinburg             = "Asia/Yekaterinburg"             // (UTC+06:00) エカテリンブルグ
	Asia_Yangon                    = "Asia/Yangon"                    // (UTC+06:30) ヤンゴン
	Asia_Bangkok                   = "Asia/Bangkok"                   // (UTC+07:00) バンコク、ハノイ、ジャカルタ
	Asia_Novosibirsk               = "Asia/Novosibirsk"               // (UTC+07:00) ノヴォシビルスク
	Asia_Krasnoyarsk               = "Asia/Krasnoyarsk"               // (UTC+08:00) クラスノヤルスク
	Asia_Ulaanbaatar               = "Asia/Ulaanbaatar"               // (UTC+08:00) ウランバートル
	Asia_Shanghai                  = "Asia/Shanghai"                  // (UTC+08:00) 北京、重慶、香港、ウルムチ
	Australia_Perth                = "Australia/Perth"                // (UTC+08:00) パース
	Asia_Singapore                 = "Asia/Singapore"                 // (UTC+08:00) クアラルンプール、シンガポール
	Asia_Taipei                    = "Asia/Taipei"                    // (UTC+08:00) 台北
	Asia_Irkutsk                   = "Asia/Irkutsk"                   // (UTC+09:00) イルクーツク
	Asia_Seoul                     = "Asia/Seoul"                     // (UTC+09:00) ソウル
	Asia_Tokyo                     = "Asia/Tokyo"                     // (UTC+09:00) 大阪、札幌、東京
	Australia_Darwin               = "Australia/Darwin"               // (UTC+09:30) ダーウィン
	Australia_Adelaide             = "Australia/Adelaide"             // (UTC+09:30) アデレード
	Australia_Hobart               = "Australia/Hobart"               // (UTC+10:00) ホバート
	Asia_Yakutsk                   = "Asia/Yakutsk"                   // (UTC+10:00) ヤクーツク
	Australia_Brisbane             = "Australia/Brisbane"             // (UTC+10:00) ブリスベン
	Pacific_Port_Moresby           = "Pacific/Port_Moresby"           // (UTC+10:00) グアム、ポートモレスビー
	Australia_Sydney               = "Australia/Sydney"               // (UTC+10:00) キャンベラ、メルボルン、シドニー
	Asia_Vladivostok               = "Asia/Vladivostok"               // (UTC+11:00) ウラジオストク
	Pacific_Guadalcanal            = "Pacific/Guadalcanal"            // (UTC+11:00) ソロモン諸島、ニューカレドニア
	Etc_GMT_M12                    = "Etc/GMT-12"                     // (UTC+12:00) 協定世界時+12
	Pacific_Fiji                   = "Pacific/Fiji"                   // (UTC+12:00) フィジー、マーシャル諸島
	Asia_Magadan                   = "Asia/Magadan"                   // (UTC+12:00) マガダン
	Pacific_Auckland               = "Pacific/Auckland"               // (UTC+12:00) オークランド、ウェリントン
	Pacific_Tongatapu              = "Pacific/Tongatapu"              // (UTC+13:00) ヌクアロファ
	Pacific_Apia                   = "Pacific/Apia"                   // (UTC+13:00) サモア
)