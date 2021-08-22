SELECT c.id, c.username, p.username as parentUserName
FROM USER c
LEFT JOIN USER p ON c.parent = p.id
ORDER BY c.id