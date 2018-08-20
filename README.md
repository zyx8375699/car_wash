# 接口

## 登录

#### Request

```
{
	"user":"user",
	"password":"password"
}
```

#### Response

返回一个token，后面的请求的header需要添加Authorization: token

```
{
    "success": true,
    "msg": "",
    "obj": "Basic token"
}
```

## 用户

### 添加用户

http POST请求

url: http://host:port/car_wash/v1/user

#### Request

```
{
    "user":{
        "name":"zyx", //姓名
        "phone":13788888888, //电话
        "type":1,  //类型
        "restMoney":1000  //余额
        "vehicleLicense": "苏J12345" //车牌
    }
}
```

#### Response

```
{
    "success": true,
    "msg": "",
    "obj": ""
}
```

### 查询用户

http GET请求

url: http://host:port/car_wash/v1/user

#### Response

```
{
    "success": true,
    "msg": "",
    "obj": [
        {
            "id": 3,
            "name": "zyx",
            "phone": 13788888888,
            "type": 1,
            "restMoney": 1000,
            "vehicleLicense": ""
        }
    ]
}
```

### 修改用户

http PUT请求

url: http://host:port/car_wash/v1/user

#### Request

```
{
	"phone":1375108989,
	"user": {
	    "name":"zyx",
	    "phone":13788888888,
	    "type":2,
	    "restMoney":3000
	}
}
```

#### Response

```
{"success":true,"msg":"","obj":null}
```

### 删除用户

http DELETE请求

url: http://host:port/car_wash/v1/user/{id}

#### Response

{"success":true,"msg":"","obj":null}

### 充值

http POST请求

url: http://host:port/car_wash/v1/user/money

#### Request

```
{
    "phone":13788889898,
    "money":1000
}
```

#### Response

```
{"success":true,"msg":"","obj":null}
```

## 交易

### 查询交易

http GET请求

url: http://host:port/car_wash/v1/transaction

#### Response

```
{
    "success": true,
    "msg": "",
    "obj": [
        {
            "id": 1, //主键
            "type": "精洗",  //交易类型
            "cost": 150,   //交易金额
            "date": "2018-08-06 15:00:00",  //交易时间
            "license": "苏J12345",  //车牌
            "phone": 13888888888,   //会员手机
            "payMethod": 2  //支付方式
        }
    ]
}
```

### 新增交易

http POST请求

url: http://host:port/car_wash/v1/transaction

#### Request

```
{
    "transaction":{
            "type": "精洗",
            "cost": 150,
            "date": "2018-08-06 15:00:00",
            "license": "苏J88099",
            "phone": 13751089898,
            "payMethod": 2
    }
}
```

#### Response

```
{"success":true,"msg":"","obj":null}
```

### 修改交易

http PUT请求

url: http://host:port/car_wash/v1/transaction

#### Request

```
{   
    "id":1,
    "transaction":{
            "type": "精洗",
            "cost": 200,
            "date": "2018-08-06 15:00:00",
            "license": "苏J88099",
            "phone": 13788888888,
            "payMethod": 2
    }
}
```

#### Response

```
{"success":true,"msg":"","obj":null}
```

### 删除交易

http DELETE请求

url: http://host:port/car_wash/v1/transaction/1

#### Response

```
{"success":true,"msg":"","obj":null}
```