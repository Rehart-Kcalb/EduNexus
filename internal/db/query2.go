package db

import "context"

const filterCourses = `-- name: FilterCourses :many
select filter.title, image,organization_name, organization_logo from filter($1,$2::bigint[]) limit $3 offset $4
`

type FilterCoursesParams struct {
	TitleParam string  `json:"title_param"`
	Column2    []int64 `json:"column_2"`
	Limit      int32   `json:"limit"`
	Offset     int32   `json:"offset"`
}

// select filter.title, image,organization_name, organization_logo from filter($1,$2::bigint[]) limit $3 offset $4;
func (q *Queries) FilterCourses(ctx context.Context, arg FilterCoursesParams) ([]GetMyCoursesRow, error) {
	rows, err := q.db.Query(ctx, filterCourses,
		arg.TitleParam,
		arg.Column2,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetMyCoursesRow
	for rows.Next() {
		var i GetMyCoursesRow
		if err := rows.Scan(
			&i.Title,
			&i.Image,
			&i.OrganizationName,
			&i.OrganizationLogo,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
