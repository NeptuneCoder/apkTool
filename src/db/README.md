将json中的配置信息放到数据中，该目录下的工具用于获取相关设计的实现



//用户表
CREATE TABLE NAME TAB_USER(
_id INT NOT NULL AUTO_INCREMENT,
   userId VARCHAR(15),
   phoneNum VARCHAR(14),
   pwd VARCHAR(60),
   isVerify VARCHAR(40),
   registerTime VARCHAR(40),
   PRIMARY KEY ( _id )
)

//游戏配置信息
CREATE TABLE NAME TAB_GAME(
    _id INT NOT NULL AUTO_INCREMENT,
    gameName VARCHAR(40),
    apktoolVersion VARCHAR(40),
    debug VARCHAR(40),
    folder VARCHAR(40),
    foyoentid VARCHAR(40),
    gameMainName VARCHAR(40),
    gameId VARCHAR(40),
    orientation VARCHAR(40)
)

{
  "channels": [
    {
      "name": "飓风",
      "icon": "rb",
      "id": "foyoent",
      "splash": true,
      "suffix": ".yehua",
      "packageName": "sg.dhxy.android.game",
      "param": [
        {
          "name": "FoyoentID",
          "value": "5008"
        },
        {
          "name": "FoyoentCID",
          "value": "1042"
        },
        {
          "name": "PAY_ACTION",
          "value": "tulongzhi.jufeng.MainActivity"
        }
      ],
      "metaData": [
        {
          "name": "CID",
          "value": "5qwan"
        },
        {
          "name": "PID",
          "value": "5qwan"
        },
        {
          "name": "APPID",
          "value": "tulongzhi"
        },
        {
          "name": "APP_CONFIG",
          "value": "data/www/config"
        }
      ]
    }
  ]
}

//渠道信息
CREATE TABLE NAME TAB_ChannelConfig(
    _id INT NOT NULL AUTO_INCREMENT,
    icon VARCHAR(40),
    gameId VARCHAR(40),
    splash VARCHAR(40),
    suffix VARCHAR(40),
    packageName VARCHAR(40),
    param VARCHAR(40),
    metaData VARCHAR(40)

)


// sdk的配置信息
CREATE TABLE NAME TAB_SdkConfig(
_id INT NOT NULL AUTO_INCREMENT,
)



































