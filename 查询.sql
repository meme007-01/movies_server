-- 获取影片信息
SELECT title,video.id as videoId,category.name as categoryName,subCategory.name as subCategoryName FROM cms.videos video,(select id,name from cms.category category) as subCategory, cms.category category 
where category_pid=category.id and category_child_id=subCategory.id;


-- 获取推荐页面数据,video.cycle=1 and  其他分类的数据查询后再动态分类,不包含禁用的分类数据
select sum,c.name,c.sort, v.id,v.title,category_pid,category_child_id,surface_plot,recommend,cycle,cycle_img,charging_mode,buy_mode,gold,directors,actors,
imdb_score,imdb_score_id,douban_score,douban_score_id,introduce,popularity_day,popularity_week,popularity_month,popularity_sum,v.note,year,album_id,v.status,v.create_at,
v.update_at,duration,region,v.language,label,v.number,v.total,horizontal_poster,vertical_poster,publish,serial_number,screenshot,gif,
alias,release_at,shelf_at,end,unit,watch,collection_id,use_local_image,titles_time,trailer_time,v.site_id,category_pid_status,category_child_id_status,play_url,play_url_put_in
from cms.videos as v 
Left join cms.category as c 
left join (select count(*) as sum,id as tId from cms.videos) tems 
on tId=v.id
on v.category_pid=c.id where v.recommend = 1 and c.status=1;


-- 查询分类下面的 banner 信息
select c.name as type_ame,c.sort as type_sort, v.id,v.title,category_pid,category_child_id,surface_plot,recommend,cycle,cycle_img,charging_mode,buy_mode,gold,directors,actors,
imdb_score,imdb_score_id,douban_score,douban_score_id,introduce,popularity_day,popularity_week,popularity_month,popularity_sum,v.note,year,album_id,v.status,v.create_at,
v.update_at,duration,region,v.language,label,v.number,v.total,horizontal_poster,vertical_poster,publish,serial_number,screenshot,gif,
alias,release_at,shelf_at,end,unit,watch,collection_id,use_local_image,titles_time,trailer_time,v.site_id,category_pid_status,category_child_id_status,play_url,play_url_put_in
from cms.videos as v Left join cms.category as c on v.category_pid=c.id where v.category_pid=1 and v.cycle = 1 and c.status=1;



select count(*) as total from cms.videos as v Left join cms.category as c on v.category_pid=c.id  where c.status=1 and category_pid=1;

-- 获取每个大分类下的 banner 图片,推荐的不包含在内
select * from cms.videos video where video.cycle=1 and video.category_pid=53 and video.recommend=2;




-- title,video.id as videoId,category.name as categoryName,subCategory.name as subName
select count(*) from cms.videos where id=25462;





-- 注册会员
INSERT INTO `cms`.`members` (`id`,`name`, `nickname`, `group_name`, `gold_tag`, `head_img`, `sex`, `birthday`, `password`, `last_login_at`, `register_type`, `create_at`, `update_at`, `is_tourists`, `channel_id`, `domain_id`, `site_id`, `member_group_id`, `register_mode`, `got_vip_days`) 
VALUES (uuid(),'test5', 'test5', '2', '1', '/data/uploadFile/00a3a5bfba1c47e0a145f7ab9e2f92d4_Su.jpeg', '3', '0', '14e1b600b1fd579f47433b88e8d85291', '0', 'web', '1726896620', '1726896620', '2', '0', '0', '1', '3', '3', '0');

-- 同时开通代理
INSERT INTO `cms`.`agent_member` (`member_id`, `brokerage_total`, `brokerage_usable`, `brokerage_freeze`, `brokerage_cashed`, `level`, `lock_level`, `performance_total`, `recharge_total`, `invite_code`, `create_at`, `update_at`) VALUES ('b1dd60a9-77e1-11ef-95d5-88a4c2a208de', '0', '0', '0', '0', '1', '0', '0', '0', '8pdMjZpc', '1726898320', '1726898320');















