package prompt

import "github.com/ollama/ollama/api"

var (
	SystemRole = api.Message{
		Role:    "system",
		Content: "You are an intent classification engine for an AI agent. You do not answer user questions.",
	}

	SystemTask = api.Message{
		Role:    "system",
		Content: `Choose exactly one intent from the allowed list. If none apply, choose "unsupported".`,
	}

	IntentTaxonomy = api.Message{
		Role: "system",
		Content: `Allowed intents and their taxonomy:

historical_lookup:
- Meaning: Retrieve specific past house price.
- Includes: results for a specific date.
- Excludes: statistics, trends, predictions.

statistical_aggregation:
- Meaning: Compute counts, frequencies, maxima, minima, or distributions acros historical lottery data.
- Includes: most frequent number, highest occurrence.
- Excludes: future predictions, single house lookups.

metadata_query:
- Meaning: Questions about dataset coverage or structure.
- Includes: number of house prices stored, date range.
- Excludes: draw results, statistics.

prediction_request:
- Meaning: Requests to forecast or guess house prices.
- Always unsupported.

general_explanation:
- Meaning: Conceptual explanations not requiring data access.
- Includes: how house are priced.

unsupported:
- Meaning: Requests outside system capability.
`,
	}
)
