-- +migrate Up
CREATE TABLE campaign
(
    id SERIAL PRIMARY KEY,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone,
    removed_at timestamp with time zone,
    frequency          text,
    account            text,
    campaign           text,
    budget             NUMERIC ,
    budget_type        text,
    campaign_type      text,
    networks           text,
    languages          text,
    bid_strategy_type  text,
    bid_strategy_name   text,
    start_date          text,
    end_date            text,
    ad_schedule         text,
    ad_rotation         text,
    inventory_type      text,
    delivery_method     text,
    targeting_method    text,
    exclusion_method    text,
    dsa_website         text,
    dsa_language        text,
    dsa_targeting_source text,
    dsa_page_feeds       text,
    flexible_reach      text
);




CREATE TABLE tech_adgroup
(
        id SERIAL PRIMARY KEY,
        campaign_id integer REFERENCES campaign (id),
        ad_group text,
        max_cpc numeric,
        max_cpm numeric,
        target_cpa text,
        max_cpv numeric,
        target_cpm numeric,
        target_roas text,
        desktop_bid_modifier text,
        mobile_bid_modifier text,
        tablet_bid_modifier text,
        tv_screen_bid_modifier text,
        top_content_bid_modifier text,
        display_network text,
        custom_bid_type text,
        targeting_optimization text,
        ad_group_type text,
        top_of_page_bid text,
        tracking_template text,
        first_page_bid text,
        final_url_suffix text,
        custom_parameters text,
        bid_modifier text,
        destination_url text,
        criterion_type text,
        reach text,
        feed text,
        radius text,
        unit text
);


CREATE TABLE adgroup
 (
    id SERIAL PRIMARY KEY,
    tech_adgroup_id integer REFERENCES tech_adgroup (id),
    folder_id integer REFERENCES folder (id),
    adgroup_id text[],
    gender text[],
    age text[],
    channel_id text[],
    audience text[],
    custom_audience text[],
    devices text[],
    placement text[],
    bls_split text[],
    topics text[],
    interest text[],
    keyword text[],
    created_at timestamp with time zone NOT NULL DEFAULT now()
 );





CREATE TABLE ad
(
    id SERIAL PRIMARY KEY,
    adgroup_id integer  REFERENCES adgroup (id),
    ad_name         text unique,
    video_id        text,
    display_url     text,
    campaign_status text,
    ad_group_status text,
    status          text,
    approval_status text,
    comment         text,
    final_url                       text,
    final_mobile_url                text,
    bumper_ad                       text,
    created_at timestamp with time zone NOT NULL DEFAULT now()

);
-- +migrate Down
DROP TABLE ad;
DROP TABLE adgroup;
DROP TABLE tech_adgroup;
DROP TABLE campaign;
DROP TABLE folder;
