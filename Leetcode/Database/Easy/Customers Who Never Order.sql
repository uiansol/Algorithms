SELECT
    name Customers
FROM
    Customers
WHERE
    id NOT IN (
        SELECT customerID FROM Orders
    )
;