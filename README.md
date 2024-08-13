# 資料庫測驗

## 題目一

```mysql
select bnb_id, bnbs.name as "bnb_name", sum(amouont) as "may_amount" 
from orders
join bnbs on (orders.bnb_id = bnbs.id)
where created_at between "2023-05-01" and "2023-05-31" and currency = "TWD"
group by bnb_id
order by may_amount desc 
limit 10
```

## 題目二


我的做法是先根據``bnb_id``進行群組，接下來進行日期和貨幣類別進行判斷，最後就是根據``sum(amouont)``
進行降序取出前十筆。

``between``和``like``選擇，我選擇用between進行2023年5月的篩選，因為like會需要逐字比對，比起between範圍查找來說會慢非常多。

條件的順序也是會影響效能的因素之一，擺在越前面的條件能將搜索的範圍縮小越小越好，在這個題目上我是先進行日期的篩選，
能將搜尋範圍縮小到2023年5月，之後再搜尋幣別為TWD。

我認為如果要再提升效率的話，可以在``created_at``和``currency``加上``index``，因為是有序的，在搜尋上可以加快搜尋效率。

# API 實作測驗

- S: 在我的程式碼中，request進來進行解析後儲存在``*model.CreateOrder``裡面，如果需要進行資料庫操作必須要再定義一個``order``進行操作，
     不能在繼續使用``*model.CreateOrder``進行資料庫操作。
- O: 假設今天``Address``其他國家可能會有一些特殊的格式，那這邊我會選擇用繼承的方式加入新的方法。
- L: 我的``CreateOrder``物件中包含了``Validate()``進行資料驗證，假設今天需要繼承``CreateOrder``，那麼子類也可以使用``Validate()``。
- I, D: 在目前的需求中，我的程式碼沒有使用到接口的應用。

在我的項目中，假設今天有一個請求進來，會先進到controller層進行request的解析和欄位型態檢查，隨後進到service層進行格式檢查和業務需求。

model層主要是放置物件相關。utils為工具包，例如``ContainNonEnglish(s)``檢查是否存在非英文字母。