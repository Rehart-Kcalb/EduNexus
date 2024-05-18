CREATE OR REPLACE FUNCTION filter(
    title_param text, 
    category_ids bigint[], 
    "limit_param" int, 
    "offset_param" int
)
RETURNS TABLE (
    title varchar(100),
    organization_name varchar(100),
    organization_logo varchar(100),
	image varchar(100)
) AS $$
BEGIN
    IF title_param = '' AND array_length(category_ids, 1) IS NULL THEN
        -- Case where both title_param and category_ids are empty
        RETURN QUERY
        SELECT
            courses.title,
            users.firstname AS organization_name,
			users.profile as organization_logo,
            courses.image
        FROM
            courses
            LEFT JOIN users ON users.id = courses.course_provider
        WHERE
            users.id = courses.course_provider
        LIMIT limit_param
        OFFSET offset_param;
    ELSE IF title_param <> '' AND array_length(category_ids, 1) IS NULL THEN
        -- Case where title_param is not empty, but category_ids is empty
        RETURN QUERY
        SELECT
            courses.title,
            users.firstname AS organization_name,
			users.profile as organization_logo,
            courses.image
        FROM
            courses
            LEFT JOIN users ON users.id = courses.course_provider
        WHERE
            courses.title ILIKE '%' || title_param || '%'
        LIMIT limit_param
        OFFSET offset_param;
    ELSE IF title_param = '' AND array_length(category_ids, 1) IS NOT NULL THEN
        -- Case where title_param is empty, but category_ids is not empty
        RETURN QUERY
        SELECT
            Distinct(courses.title),
            users.firstname AS organization_name,
			users.profile as organization_logo,
            courses.image
        FROM
            courses
            LEFT JOIN users ON users.id = courses.course_provider
            Left join course_categories on courses.id = course_categories.course_id
        WHERE
            course_categories.category_id = ANY(category_ids)
        LIMIT limit_param
        OFFSET offset_param;
    ELSE
        -- Case where both title_param and category_ids are not empty
        RETURN QUERY
        SELECT
            Distinct(courses.title),
            users.firstname AS organization_name,
			users.profile as organization_logo,
            courses.image
        FROM
            courses
            LEFT JOIN users ON users.id = courses.course_provider
            Left join course_categories on courses.id = course_categories.course_id
        WHERE
            courses.title ILIKE '%' || title_param || '%'
            AND course_categories.category_id = ANY(category_ids)
        LIMIT limit_param
        OFFSET offset_param;
	END IF;
	END IF;
    END IF;
END;
$$ LANGUAGE plpgsql;
