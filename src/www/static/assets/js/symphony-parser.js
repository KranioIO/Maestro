function generate_symphony_nodes(json) {
    let nodes = JSON.parse(json);
    let cnodes = nodes.map(node => {return {data: {id: node.name}}});

    return cnodes;
}

function generate_symphony_edges(json) {
    let nodes = JSON.parse(json);
    let edges = []

    nodes.map(node => {
        if (!node.depends)
            return;

        node.depends.map(dependency => {
            edges.push({data: {id: `${dependency}->${node.name}`, weight: 2, source: dependency, target: node.name}})
        })
    });

    return edges;
}

function set_node_bg (node) {
    // set color based on status
    // return '#19aec6';

    return '#59bdcc #3cb1c3';
}

// https://js.cytoscape.org/#style

var cytoscape_style = [
    {
        selector: 'node',
        css: {
            'shape'         : 'round-rectangle',
            'content'       : 'data(id)',
            'text-valign'   : 'center',
            'text-halign'   : 'center',
            'height'        : 'label',
            'width'         : 'label',
            'padding'       : '5',
            'color'         : 'white',
            'font-size'     : '10px',
            'font-family'   : 'Roboto',
            'font-weight'   : 300,

            'shadow-blur': 20,
            'shadow-color': 'red',
            'shadow-offset-x': 10,
            'shadow-offset-y': 10,
            'shadow-opacity': 1,

            // 'background-fill': 'linear-gradient',
            // 'background-gradient-stop-colors': set_node_bg,
            // 'background-gradient-direction': 'to-left',
            'background-color': '#00bcd4',
        }
    },
    {
        selector: 'edge',
        css: {
            'target-arrow-shape'        : 'triangle',
            "curve-style"               : "taxi",
            "taxi-direction"            : "downward",
            "taxi-turn-min-distance"    : 15,
            "taxi-turn"                 : 15,
            'width'                     : 1,
        }
    }
]