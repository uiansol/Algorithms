SELECT
    u.user_id buyer_id, 
    u.join_date, 
    count(o.order_id) orders_in_2019
FROM
    Users u LEFT JOIN ORDERS o ON u.user_id = o.buyer_id AND YEAR(o.order_date)='2019'
GROUP BY
    u.user_id
;