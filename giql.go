package giql

import (
	"html/template"
	"net/http"
)

func New(endpoint ...string) http.HandlerFunc {
	t := template.New("graphiql")
	t, err := t.Parse(graphiqlHtml)
	if err != nil {
		panic(err)
	}
	ep := template.JS("window.location.origin + '/graphql'")
	if len(endpoint) > 0 && endpoint[0] != "" {
		ep = template.JS("'" + endpoint[0] + "'")
	}
	return func(w http.ResponseWriter, r *http.Request) {
		t.Execute(w, map[string]interface{}{
			"Endpoint": ep,
		})
	}
}

const graphiqlHtml = `
<!--
 *  Copyright (c) 2021 GraphQL Contributors
 *  All rights reserved.
 *
 *  This source code is licensed under the license found in the
 *  LICENSE file in the root directory of this source tree.
-->
<!DOCTYPE html>
<html>
<head>
    <style>
        body {
            height: 100%;
            margin: 0;
            width: 100%;
            overflow: hidden;
        }

        #graphiql {
            height: 100vh;
        }
    </style>
    <script crossorigin src="https://unpkg.com/react/umd/react.production.min.js"></script>
    <script crossorigin src="https://unpkg.com/react-dom/umd/react-dom.production.min.js"></script>
    <link rel="stylesheet" href="https://unpkg.com/graphiql/graphiql.min.css"/>
</head>

<body>
<div id="graphiql">Loading...</div>
<script src="https://unpkg.com/graphiql/graphiql.min.js" type="application/javascript"></script>
<script>
  function graphQLFetcher(graphQLParams) {
    return fetch(
      {{ .Endpoint }} + window.location.search,
      {
        method: 'post',
        headers: {
          Accept: 'application/json',
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(graphQLParams),
      },
    ).then(function (response) {
      return response.json().catch(function () {
        return response.text();
      });
    });
  }

  ReactDOM.render(
    React.createElement(GraphiQL, {
      fetcher: graphQLFetcher,
      defaultVariableEditorOpen: true,
    }),
    document.getElementById('graphiql'),
  );
</script>
</body>
</html>
`
