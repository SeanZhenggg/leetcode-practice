SELECT e3.salary as SecondHighestSalary
FROM (SELECT e2.salary as salary
      FROM Employee as e1
               LEFT JOIN Employee as e2 ON e1.salary > e2.salary AND e1.id != e2.id
      ORDER BY e1.salary DESC, e2.salary DESC LIMIT 1) as e3;