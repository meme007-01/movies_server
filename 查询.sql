-- 获取影片信息
SELECT title,video.id as videoId,category.name as categoryName,subCategory.name as subCategoryName FROM cms.videos video,(select id,name from cms.category category) as subCategory, cms.category category 
where category_pid=category.id and category_child_id=subCategory.id;


-- 获取推荐页面数据,video.cycle=1 and  其他分类的数据查询后再动态分类,不包含禁用的分类数据
select c.name,c.sort, v.id,v.title,category_pid,category_child_id,surface_plot,recommend,cycle,cycle_img,charging_mode,buy_mode,gold,directors,actors,
imdb_score,imdb_score_id,douban_score,douban_score_id,introduce,popularity_day,popularity_week,popularity_month,popularity_sum,v.note,year,album_id,v.status,v.create_at,
v.update_at,duration,region,v.language,label,v.number,v.total,horizontal_poster,vertical_poster,publish,serial_number,screenshot,gif,
alias,release_at,shelf_at,end,unit,watch,collection_id,use_local_image,titles_time,trailer_time,v.site_id,category_pid_status,category_child_id_status,play_url,play_url_put_in
from cms.videos as v Left join cms.category as c on v.category_pid=c.id where v.recommend = 1;


-- 获取每个大分类下的 banner 图片,推荐的不包含在内
select * from cms.videos video where video.cycle=1 and video.category_pid=53 and video.recommend=2;






-- title,video.id as videoId,category.name as categoryName,subCategory.name as subName
select count(*) from cms.videos where id=25462;