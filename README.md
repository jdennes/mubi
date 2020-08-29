# mubi

Do things with https://mubi.com.

Example usage:

```go
package main

import (
	"fmt"
	"github.com/jdennes/mubi"
)

func main() {
	api := mubi.NewMubiAPI()
	userId, page, perPage := int64(7995037), 1, 20
	ratings := api.GetRatings(userId, page, perPage)

	for _, rating := range ratings {
		fmt.Printf("%s (%d) - Rated %d/5\n", rating.Film.Title, rating.Film.Year, rating.Overall)
	}
}
```

Example command provided:

```
$ mubi ratings --userid=7995037 --page=1 --per-page=5
Dear Zachary (2008) - http://mubi.com/films/dear-zachary
Directed by Kurt Kuenne
Rated 4/5 stars on 2020-08-28 23:47:58 +0200 CEST
-----
Papicha (2019) - http://mubi.com/films/papicha
Directed by Mounia Meddour Gens
Rated 5/5 stars on 2020-08-26 22:15:02 +0200 CEST
-----
Leviathan (2014) - http://mubi.com/films/leviathan-2014
Directed by Andrey Zvyagintsev
Rated 5/5 stars on 2020-08-19 18:47:59 +0200 CEST
-----
Amy (2015) - http://mubi.com/films/amy-2015
Directed by Asif Kapadia
Rated 4/5 stars on 2020-08-18 11:23:48 +0200 CEST
-----
Hoop Dreams (1994) - http://mubi.com/films/hoop-dreams
Directed by Steve James
Rated 4/5 stars on 2020-08-18 10:58:18 +0200 CEST
-----
```

```
$ mubi ratings --userid=7995037 --page=1 --per-page=5 --export-for-letterboxd
Title,Year,Directors,Rating,WatchedDate
Dear Zachary,2008,Kurt Kuenne,4,2020-08-28
Papicha,2019,Mounia Meddour Gens,5,2020-08-26
Leviathan,2014,Andrey Zvyagintsev,5,2020-08-19
Amy,2015,Asif Kapadia,4,2020-08-18
Hoop Dreams,1994,Steve James,4,2020-08-18
```
