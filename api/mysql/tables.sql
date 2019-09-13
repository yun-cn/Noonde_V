
-- ================================================================================
-- users
-- --------------------------------------------------------------------------------

drop   table if exists users;
create table           users (
    id              int          unsigned not null primary key auto_increment,
    email           varchar(191)          not null default '',
    password        text                  not null,
    nickname        varchar(191)          not null default '',
    avatar          varchar(191)          not null default '',
    avatar_sum      varchar(191)          not null default '',
    profile         text                  not null,
    max_tokens      tinyint      unsigned not null default 5,
    failed          tinyint      unsigned not null default 0,
    created         datetime              not null default current_timestamp,
    updated         datetime              not null default current_timestamp
);
alter table users auto_increment=10001;
create unique index users_i01 on users (email);


-- ================================================================================
-- user_tokens
-- --------------------------------------------------------------------------------

drop   table if exists user_tokens;
create table           user_tokens (
    id      int        unsigned      not null primary key auto_increment,
    user_id int        unsigned      not null default 0,
    token   text                     not null,
    expired  datetime                not null,
    created  datetime                not null default current_timestamp,
    updated  datetime                not null default current_timestamp
);
create index user_tokens_i01 on user_tokens (user_id, expired);

# -- ================================================================================
# -- user_likes
# -- --------------------------------------------------------------------------------
#
# drop      table  if exists user_likes;
# create    table            user_likes (
#   id                int        unsigned      not null primary key auto_increment,
#   user_id           int        unsigned      not null default 0,
#   liked_note_ids    text                     not null default '',
#   liked_comment_ids text                     not null default '',
#   created           datetime                 not null default current_timestamp,
#   updated           datetime                 not null default current_timestamp
# );
# create  index   user_likes_i01  on user_likes (user_id);

-- ================================================================================
-- relationships
-- --------------------------------------------------------------------------------
drop     table  if    exists relationships;
create   table               relationships (
    id              int        unsigned      not null primary key auto_increment,
    follower_id     int        unsigned      not null default 0,
    followed_id     int        unsigned      not null default 0,
    created         datetime                 not null default current_timestamp,
    updated         datetime                 not null default current_timestamp
);
create           index  relationships_i01 on relationships (follower_id);
create           index  relationships_i02 on relationships (followed_id);
create   unique  index  relationships_i03 on relationships (followed_id, follower_id);

-- ================================================================================
-- notes
-- --------------------------------------------------------------------------------

drop   table if exists notes;
create table           notes (
    id                 int        unsigned      not null primary key auto_increment,
    uniq_key           varchar(191)             not null default '',
    user_id            int        unsigned      not null default 0,
    name               varchar(191)             not null default '',
    tags               text                     not null ,
    tweet_text         text                     not null ,
    body               text                     not null ,
    like_count         int                      not null,
    comment_count      int                      not null default 0,
    comment_viewable   tinyint                  not null default 0,
    states             tinyint    unsigned      not null default 0,
    display_date       datetime                 not null,
    can_read           tinyint                  unsigned not null default 0,
    created            datetime                 not null default current_timestamp,
    updated            datetime                 not null default current_timestamp
);
create index notes_i01    on notes (user_id, name);
alter  table notes        auto_increment = 10001;

-- ================================================================================
-- note_likes
-- --------------------------------------------------------------------------------

drop     table   if  exists note_likes;
create   table              note_likes (
    id              int        unsigned      not null primary key auto_increment,
    note_id         int        unsigned      not null default 0,
    user_id         int        unsigned      not null default 0,
    states          tinyint    unsigned      not null default 0,
    created         datetime                 not null default current_timestamp,
    updated         datetime                 not null default current_timestamp
);
create        index note_likes_i01 on note_likes (user_id);
create        index note_likes_i02 on note_likes (note_id);
create unique index note_likes_i03 on note_likes (user_id, note_id);

-- ================================================================================
-- comments
-- --------------------------------------------------------------------------------

drop       table  if exists comments;
create     table            comments (
    id              int        unsigned      not null primary key auto_increment,
    user_id         int        unsigned      not null default 0,
    note_id         int        unsigned      not null default 0,
    comment         text                     not null,
    like_count      int                      not null default 0,
    created         datetime                 not null default current_timestamp,
    updated         datetime                 not null default current_timestamp
);
create         index  comment_i01   on  comments (user_id);
create         index  comment_i02   on  comments (note_id);


-- ================================================================================
-- comment_likes
-- --------------------------------------------------------------------------------

drop      table   if exists comment_likes;
create    table             comment_likes (
    id              int        unsigned      not null primary key auto_increment,
    comment_id      int        unsigned      not null default 0,
    user_id         int        unsigned      not null default 0,
    states          tinyint    unsigned      not null default 0,
    created         datetime                 not null default current_timestamp,
    updated         datetime                 not null default current_timestamp
);
create         index  comment_likes_i01   on  comment_likes (user_id);
create         index  comment_likes_i02   on  comment_likes (comment_id);
create unique  index  comment_likes_i03   on  comment_likes (user_id, comment_id);

-- ================================================================================
-- spacemarket_event_types
-- --------------------------------------------------------------------------------

drop      table   if exists spacemarket_event_types;
create    table             spacemarket_event_types (
    id                 int                 unsigned      not null primary key auto_increment,
    state              int                 unsigned      not null default 1,
    event_type         varchar(191)                      not null default '',
    event_type_text    varchar(191)                      not null default '',
    start_page         int                 unsigned      not null default 1,
    hourly_at          datetime                                            ,
    daily_at           datetime                                            ,
    created            datetime                          not null default current_timestamp,
    updated            datetime                          not null default current_timestamp
);
create   index        spacemarket_event_types_i01 on  spacemarket_event_types(event_type);
insert   into         spacemarket_event_types (event_type, event_type_text) value ('party', 'パーティー');
insert   into         spacemarket_event_types (event_type, event_type_text) value ('class', '会議・研修');
insert   into         spacemarket_event_types (event_type, event_type_text) value ('photo_shoot', '写真撮影');
insert   into         spacemarket_event_types (event_type, event_type_text) value ('film_shoot', 'ロケ撮影');
insert   into         spacemarket_event_types (event_type, event_type_text) value ('social_event', 'イベント');
insert   into         spacemarket_event_types (event_type, event_type_text) value ('performance', '演奏・パフォーマンス');
insert   into         spacemarket_event_types (event_type, event_type_text) value ('studio', '個展・展示会');
insert   into         spacemarket_event_types (event_type, event_type_text) value ('sports', 'スポーツ・フィットネス');
insert   into         spacemarket_event_types (event_type, event_type_text) value ('office', 'オフィス');
insert   into         spacemarket_event_types (event_type, event_type_text) value ('wedding', '結婚式');
insert   into         spacemarket_event_types (event_type, event_type_text) value ('other', 'その他');
insert   into         spacemarket_event_types (event_type, event_type_text) value ('stay_business', '出張・ビジネス');
insert   into         spacemarket_event_types (event_type, event_type_text) value ('stay_party', 'パーティー');
insert   into         spacemarket_event_types (event_type, event_type_text) value ('stay_trip', '旅行');
insert   into         spacemarket_event_types (event_type, event_type_text) value ('stay_group', '合宿・グループ');
insert   into         spacemarket_event_types (event_type, event_type_text) value ('stay_vacation', 'バケーションレンタル');


-- ================================================================================
-- instabase_event_types
-- --------------------------------------------------------------------------------

drop      table   if exists instabase_event_types;
create    table             instabase_event_types (
   id                   int                 unsigned      not null primary key auto_increment,
   state                int                 unsigned      not null default 1,
   event_type_text      varchar(191)                      not null default '',
   event_type_en        varchar(191)                      not null default '',
   event_type           int                 unsigned      not null default 0,
   start_page           int                 unsigned      not null default 1,
   hourly_at            datetime                                            ,
   created              datetime                          not null default current_timestamp,
   updated              datetime                          not null default current_timestamp
);
create      index       instabase_event_types_i01 on instabase_event_types(event_type_text);
create      index       instabase_event_types_i02 on instabase_event_types(event_type_en);
insert      into        instabase_event_types        (event_type_text, event_type_en ,event_type) values ('打ち合わせ・商談', 'meeting' ,1);
insert      into        instabase_event_types        (event_type_text, event_type_en ,event_type) values ('セミナー・研修', 'seminar', 3);
insert      into        instabase_event_types        (event_type_text, event_type_en ,event_type) values ('レッスン・講座', 'lesson', 2);
insert      into        instabase_event_types        (event_type_text, event_type_en ,event_type) values ('ヨガ・ダンス', 'studio' ,7);
insert      into        instabase_event_types        (event_type_text, event_type_en ,event_type) values ('パーティー', 'party', 4);
insert      into        instabase_event_types        (event_type_text, event_type_en ,event_type) values ('女子会・ママ会', 'girls-party', 14);
insert      into        instabase_event_types        (event_type_text, event_type_en ,event_type) values ('撮影・収録', 'photo-studio', 10);
insert      into        instabase_event_types        (event_type_text, event_type_en ,event_type) values ('面接・試験', 'interview', 11);
insert      into        instabase_event_types        (event_type_text, event_type_en ,event_type) values ('セラピー', 'treatment', 6);
insert      into        instabase_event_types        (event_type_text, event_type_en ,event_type) values ('ワークショップ', 'workshop', 15);
insert      into        instabase_event_types        (event_type_text, event_type_en ,event_type) values ('カウンセリング', 'therapy', 13);
insert      into        instabase_event_types        (event_type_text, event_type_en ,event_type) values ('上映会・映画鑑賞', 'screening', 16);
insert      into        instabase_event_types        (event_type_text, event_type_en ,event_type) values ('勉強会', 'study', 5);
insert      into        instabase_event_types        (event_type_text, event_type_en ,event_type) values ('作業場所', 'desk' ,8);
insert      into        instabase_event_types        (event_type_text, event_type_en ,event_type) values ('ボードゲーム', 'board-game', 12);


-- ================================================================================
-- platform
-- --------------------------------------------------------------------------------
drop      table   if exists platforms;
create    table             platforms (
    id                  int                 unsigned      not null primary key auto_increment,
    url                 varchar(191)                      not null default '',
    listing_url         varchar(191)                      not null default '',
    created             datetime                          not null default current_timestamp,
    updated             datetime                          not null default current_timestamp
);
create    index         platforms_i01 on platforms (listing_url);
insert    into          platforms   (url,listing_url)  values ('https://www.spacemarket.com', 'https://www.spacemarket.com/spaces/');
insert    into          platforms   (url,listing_url)  values ('https://www.instabase.jp', 'https://www.instabase.jp/space/');


-- ================================================================================
-- Spaces
-- --------------------------------------------------------------------------------

drop      table   if exists spaces;
create    table             spaces (
    id                       int                 unsigned      not null primary key auto_increment,
    iop                      int                               not null default 0,
    uip                      varchar(191)                      not null default '',
    hiop                     int                               not null default 0,
    name                     varchar(191)                      not null default '',
    capacity                 int                               not null default 1,
    description              text                              not null,
    equipment_description    text                              not null,
    amenities                text                              not null,
    event_types              text                              not null,
    hourly_min_price         decimal(10,2)                     not null default 0.0,
    hourly_max_price         decimal(10,2)                     not null default 0.0,
    daily_min_price          decimal(10,2)                     not null default 0.0,
    daily_max_price          decimal(10,2)                     not null default 0.0,
    embed_video_url          varchar(191)                      not null default '',
    state_text               varchar(191)                      not null default '',
    city                     varchar(191)                      not null default '',
    address                  varchar(191)                      not null default '',
    latitude                 decimal(18,15)                    not null default 0.0,
    longitude                decimal(18,15)                    not null default 0.0,
    access                   text                              ,
    thumbnails               text                              ,
    third_review_score       decimal(10,8)                     not null default 0.0,
    third_review_count       int                               not null default 0,
    third_reply_rate         decimal(6,3)                      not null default 0.0,
    platform_id              int                               not null default 0,
    hash                     varchar(191)                      not null default '',
    hash_at                  datetime,
    space_url                varchar(191)                      not null default '',
    states                   int                  unsigned     not null default 1,
    created                  datetime                          not null default current_timestamp,
    updated                  datetime                          not null default current_timestamp
);
alter                        table          spaces             auto_increment = 10001;
create unique                index          spaces_i01                        on spaces       (iop, platform_id);
create                       index          spaces_i02                        on spaces       (iop);
create                       index          spaces_i03                        on spaces       (uip);
create                       index          spaces_i04                        on spaces       (hiop);
create                       index          spaces_i05                        on spaces       (name);
create                       index          spaces_i06                        on spaces       (hourly_min_price);
create                       index          spaces_i07                        on spaces       (hourly_max_price);


-- ================================================================================
-- space_calendars
-- --------------------------------------------------------------------------------

-- drop         table  if exists space_calendars;
-- create       table            space_calendars(
--     id                       int                 unsigned      not null primary key auto_increment,
--     space_id                 int                               not null default 0,
--     year                     int                 unsigned      not null default 1970,
--     month                    int                 unsigned      not null default 0,
--     day                      int                 unsigned      not null default 0,
--     hour                     int                 unsigned      not null default 0,
--     state                    int                 unsigned      not null default 1,
--     price                    decimal(10,2)                     not null default 0.0
-- )
-- create unique                index         space_calendars_i01             on space_calendars  (space_id, year, month, day, hour);


