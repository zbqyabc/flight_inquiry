## 请求地址
http://localhost:8080/flight/get_new_list

请求参数
```请求参数
{
    "flight_list":[
        ["a","b"],
        ["e","f"],
        ["c","d"],
        ["b","c"],
        ["d","e"]
		
    ]
}
```
返回参数
```返回参数
{
    "arrFlight": [
        "a",
        "f"
    ],
    "code": 200,
    "msg": "Success",
    "newList": [
        [
            "a",
            "b"
        ],
        [
            "b",
            "c"
        ],
        [
            "c",
            "d"
        ],
        [
            "d",
            "e"
        ],
        [
            "e",
            "f"
        ]
    ]
}
```
