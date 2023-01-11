-- Remove the 'web' user and their privileges
DROP USER 'web'@'%';

-- Remove the dummy records 
DELETE FROM snippets WHERE title = 'An old silent pond';
DELETE FROM snippets WHERE title = 'Over the wintry forest';
DELETE FROM snippets WHERE title = 'First autumn morning';

-- Remove the index
DROP INDEX idx_snippets_created ON snippets;

-- Remove the 'snippets' table
DROP TABLE snippets;