-- 获取影片信息
SELECT title,video.id as videoId,category.name as categoryName,subCategory.name as subCategoryName FROM cms.videos video,(select id,name from cms.category category) as subCategory, cms.category category 
where category_pid=category.id and category_child_id=subCategory.id;


-- 获取推荐页面数据,video.cycle=1 and  其他分类的数据查询后再动态分类
select * from cms.videos video where  video.recommend=1;


-- 获取每个大分类下的 banner 图片,推荐的不包含在内
select * from cms.videos video where video.cycle=1 and video.category_pid=53 and video.recommend=2;






-- title,video.id as videoId,category.name as categoryName,subCategory.name as subName
select count(*) from cms.videos where id=25462;