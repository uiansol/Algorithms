SELECT
    *
FROM
(SELECT
    employee_id
FROM 
    Employees e
WHERE
    employee_id NOT IN (
        SELECT employee_id FROM Salaries
    )
) AS t1
UNION
(SELECT
    employee_id
FROM 
    Salaries s
WHERE
    employee_id NOT IN (
        SELECT employee_id FROM Employees
    )
)
ORDER BY
    employee_id
;
