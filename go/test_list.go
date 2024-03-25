package main

import (
    "main/list"
)

type testData struct {
    value int;
    content [600]byte;
}

// This is a random pattern of actions
var actions = [][2]int { {0, 0}, {0, 0}, {0, 0}, {0, 0}, {1, 0}, {1, 0}, {1, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {2, 1}, {2, 2}, {1, 0}, {1, 0}, {1, 0}, {1, 0}, {0, 0}, {0, 0}, {2, 5}, {0, 0}, {1, 0}, {0, 0}, {2, 7}, {2, 4}, {1, 0}, {0, 0}, {0, 0}, {1, 0}, {0, 0}, {0, 0}, {2, 3}, {0, 0}, {2, 6}, {0, 0}, {2, 9}, {1, 0}, {2, 8}, {0, 0}, {0, 0}, {0, 0}, {1, 0}, {0, 0}, {0, 0}, {1, 0}, {1, 0}, {2, 10}, {0, 0}, {1, 0}, {2, 11}, {1, 0}, {0, 0}, {2, 12}, {2, 13}, {0, 0}, {0, 0}, {0, 0}, {2, 14}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {2, 15}, {1, 0}, {1, 0}, {2, 16}, {0, 0}, {1, 0}, {0, 0}, {0, 0}, {2, 17}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {1, 0}, {1, 0}, {0, 0}, {1, 0}, {2, 18}, {2, 19}, {1, 0}, {2, 20}, {0, 0}, {0, 0}, {0, 0}, {1, 0}, {0, 0}, {1, 0}, {2, 21}, {1, 0}, {0, 0}, {2, 22}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {1, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {2, 23}, {0, 0}, {0, 0}, {0, 0}, {1, 0}, {0, 0}, {0, 0}, {0, 0}, {2, 24}, {0, 0}, {2, 25}, {2, 26}, {1, 0}, {1, 0}, {0, 0}, {0, 0}, {0, 0}, {1, 0}, {0, 0}, {2, 27}, {2, 28}, {0, 0}, {1, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {2, 29}, {0, 0}, {0, 0}, {2, 30}, {2, 31}, {1, 0}, {2, 32}, {0, 0}, {1, 0}, {1, 0}, {0, 0}, {2, 33}, {2, 0}, {2, 34}, {0, 0}, {0, 0}, {0, 0}, {1, 0}, {0, 0}, {2, 35}, {1, 0}, {0, 0}, {2, 36}, {1, 0}, {1, 0}, {0, 0}, {0, 0}, {1, 0}, {2, 37}, {0, 0}, {2, 38}, {2, 39}, {0, 0}, {1, 0}, {0, 0}, {2, 40}, {1, 0}, {0, 0}, {0, 0}, {1, 0}, {1, 0}, {1, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {1, 0}, {1, 0}, {2, 41}, {0, 0}, {2, 42}, {0, 0}, {1, 0}, {0, 0}, {2, 43}, {0, 0}, {0, 0}, {2, 44}, {0, 0}, {1, 0}, {0, 0}, {1, 0}, {0, 0}, {0, 0}, {1, 0}, {0, 0}, {0, 0}, {0, 0}, {2, 45}, {1, 0}, {0, 0}, {0, 0}, {1, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {1, 0}, {0, 0}, {0, 0}, {1, 0}, {0, 0}, {2, 46}, {2, 50}, {1, 0}, {0, 0}, {1, 0}, {0, 0}, {0, 0}, {2, 47}, {2, 48}, {1, 0}, {2, 49}, {2, 51}, {0, 0}, {0, 0}, {2, 52}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {1, 0}, {2, 57}, {0, 0}, {0, 0}, {2, 53}, {0, 0}, {0, 0}, {2, 54}, {0, 0}, {2, 55}, {0, 0}, {1, 0}, {0, 0}, {0, 0}, {2, 56}, {1, 0}, {2, 58}, {0, 0}, {0, 0}, {2, 60}, {0, 0} };

var list = llist.Head[testData]{};

func add_node(value int) *llist.Node[testData] {
    n := llist.Node[testData]{Data: testData {value: value}};
    list.Add(&n);
    return &n;
}

func test(times int) {
    var stash [256]*llist.Node[testData];

    for round := 0; round < times; round++ {
        stash_i := 0;
        for _, action := range actions {
            t, index := action[0], action[1];
            switch t {
            case 0:
                if list.IsEmpty() {
                    continue;
                }
            case 1:
                stash[stash_i] = add_node(round); stash_i++;
            case 2:
                list.Del(stash[index]);
            }
        }
    }
}

func main() {
    list.Init();
    test(1_000_000);

    /*
    // Print the values
    list.Each()(func(p *testData) {
        fmt.Printf("v: %d\n", p.value);
    });

    // Go 1.23
    for p := range list.Each() {
        fmt.Printf("v: %d\n", p.value);
    }
    */
}