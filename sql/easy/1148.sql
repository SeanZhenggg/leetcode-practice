-- solution 1 - group by
select author_id as id
from Views
group by author_id, viewer_id
having author_id = viewer_id
order by id ASC;

-- solution 2 - where + distinct
select DISTINCT author_id,
from Views
where author_id = viewer_id
order by id ASC;