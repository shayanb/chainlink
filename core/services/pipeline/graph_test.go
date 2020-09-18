package pipeline

import (
	"net/url"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
	"gonum.org/v1/gonum/graph/encoding/dot"

	"github.com/smartcontractkit/chainlink/core/store/models"
)

const dotStr = `
    // data source 1
    ds1          [type=bridge name=voter_turnout];
    ds1_parse    [type=jsonparse path="one,two"];
    ds1_multiply [type=multiply times=1.23];

    // data source 2
    ds2          [type=http method=GET url="https://chain.link/voter_turnout/USA-2020" requestData="{\"hi\": \"hello\"}"];
    ds2_parse    [type=jsonparse path="three,four"];
    ds2_multiply [type=multiply times=4.56];

    answer1 [type=median];

    ds1 -> ds1_parse -> ds1_multiply -> answer1;
    ds2 -> ds2_parse -> ds2_multiply -> answer1;

    answer2 [type=bridge name=election_winner];
`

func TestGraph_Decode(t *testing.T) {
	expected := map[string]map[string]bool{
		"ds1": {
			"ds1":          false,
			"ds1_parse":    true,
			"ds1_multiply": false,
			"ds2":          false,
			"ds2_parse":    false,
			"ds2_multiply": false,
			"answer1":      false,
			"answer2":      false,
		},
		"ds1_parse": {
			"ds1":          false,
			"ds1_parse":    false,
			"ds1_multiply": true,
			"ds2":          false,
			"ds2_parse":    false,
			"ds2_multiply": false,
			"answer1":      false,
			"answer2":      false,
		},
		"ds1_multiply": {
			"ds1":          false,
			"ds1_parse":    false,
			"ds1_multiply": false,
			"ds2":          false,
			"ds2_parse":    false,
			"ds2_multiply": false,
			"answer1":      true,
			"answer2":      false,
		},
		"ds2": {
			"ds1":          false,
			"ds1_parse":    false,
			"ds1_multiply": false,
			"ds2":          false,
			"ds2_parse":    true,
			"ds2_multiply": false,
			"answer1":      false,
			"answer2":      false,
		},
		"ds2_parse": {
			"ds1":          false,
			"ds1_parse":    false,
			"ds1_multiply": false,
			"ds2":          false,
			"ds2_parse":    false,
			"ds2_multiply": true,
			"answer1":      false,
			"answer2":      false,
		},
		"ds2_multiply": {
			"ds1":          false,
			"ds1_parse":    false,
			"ds1_multiply": false,
			"ds2":          false,
			"ds2_parse":    false,
			"ds2_multiply": false,
			"answer1":      true,
			"answer2":      false,
		},
		"answer1": {
			"ds1":          false,
			"ds1_parse":    false,
			"ds1_multiply": false,
			"ds2":          false,
			"ds2_parse":    false,
			"ds2_multiply": false,
			"answer1":      false,
			"answer2":      false,
		},
		"answer2": {
			"ds1":          false,
			"ds1_parse":    false,
			"ds1_multiply": false,
			"ds2":          false,
			"ds2_parse":    false,
			"ds2_multiply": false,
			"answer1":      false,
			"answer2":      false,
		},
	}

	g := NewTaskDAG()
	err := g.UnmarshalText([]byte(dotStr))
	require.NoError(t, err)

	nodes := make(map[string]int64)
	iter := g.Nodes()
	for iter.Next() {
		n := iter.Node().(*taskDAGNode)
		nodes[n.DOTID()] = n.ID()
	}

	for from, connections := range expected {
		for to, connected := range connections {
			require.Equal(t, connected, g.HasEdgeFromTo(nodes[from], nodes[to]))
		}
	}
}

func TestGraph_TasksInDependencyOrder(t *testing.T) {
	g := NewTaskDAG()
	err := g.UnmarshalText([]byte(dotStr))
	require.NoError(t, err)

	u, err := url.Parse("https://chain.link/voter_turnout/USA-2020")
	require.NoError(t, err)

	answer1 := &MedianTask{}
	answer2 := &BridgeTask{Name: "election_winner"}
	ds1_multiply := &MultiplyTask{
		Times:    decimal.NewFromFloat(1.23),
		BaseTask: BaseTask{outputTask: answer1},
	}
	ds1_parse := &JSONParseTask{
		Path:     []string{"one", "two"},
		BaseTask: BaseTask{outputTask: ds1_multiply},
	}
	ds1 := &BridgeTask{
		Name:     "voter_turnout",
		BaseTask: BaseTask{outputTask: ds1_parse},
	}
	ds2_multiply := &MultiplyTask{
		Times:    decimal.NewFromFloat(4.56),
		BaseTask: BaseTask{outputTask: answer1},
	}
	ds2_parse := &JSONParseTask{
		Path:     []string{"three", "four"},
		BaseTask: BaseTask{outputTask: ds2_multiply},
	}
	ds2 := &HTTPTask{
		URL:         models.WebURL(*u),
		Method:      "GET",
		RequestData: HttpRequestData{"hi": "hello"},
		BaseTask:    BaseTask{outputTask: ds2_parse},
	}

	tasks, err := g.TasksInDependencyOrder()
	require.NoError(t, err)

	// Make sure that no task appears in the array until its output task has already appeared
	for i, task := range tasks {
		if task.OutputTask() != nil {
			require.Contains(t, tasks[:i], task.OutputTask())
		}
	}

	expected := []Task{ds1, ds1_parse, ds1_multiply, ds2, ds2_parse, ds2_multiply, answer1, answer2}
	require.Len(t, tasks, len(expected))

	for _, task := range expected {
		require.Contains(t, tasks, task)
	}
}

func TestGraph_HasCycles(t *testing.T) {
	g := NewTaskDAG()
	err := g.UnmarshalText([]byte(dotStr))
	require.NoError(t, err)
	require.False(t, g.HasCycles())

	g = NewTaskDAG()
	err = dot.Unmarshal([]byte(`
        digraph {
            a [type=bridge];
            b [type=multiply times=1.23];
            a -> b -> a;
        }
    `), g)
	require.NoError(t, err)
	require.True(t, g.HasCycles())
}
