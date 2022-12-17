<template>
    <el-card class="box-card ezy-map" shadow="never">
        <div class="ezy-map-container">
            <VueFlow :apply-default="false" @dragover="onDragOver" />
        </div>
    </el-card>
</template>

<script>
export default {
    name: 'EzyMap'
}
</script>

<script setup>
import { VueFlow, useVueFlow } from '@vue-flow/core'
import { reactive, ref } from 'vue'


const defaultX = ref(200)
const defaultY = ref(5)

const stagemap = ref(
    [
        {
            "label": "A",
            "number": "12"
        },
        {
            "label": "B",
            "number": "12"
        },
        {
            "label": "C",
            "number": "12"
        },
        {
            "label": "D",
            "number": "12"
        },
        {
            "label": "E",
            "number": "10"
        },
        {
            "label": "F",
            "number": "10"
        },
        {
            "label": "G",
            "number": "10"
        },
        {
            "label": "I",
            "number": "10"
        }
    ]
)

const exampleData = ref(
    [
        {
            "area": "Khu 1",
            "seats": [
                "A1",
                "A2",
                "A3",
                "A4",
                "A5",
                "A6",
                "A7",
                "A8",
                "A9",
                "A10",
                "A11",
                "A12",
                "B1",
                "B2",
                "B3",
                "B4",
                "B5",
                "B7",
                "B6",
                "B8",
                "B9",
                "B10",
                "B11",
                "B12",
                "C1",
                "C2",
                "C3",
                "C4",
                "C5",
                "C6",
                "C7",
                "C8",
                "C9",
                "C10",
                "C11",
                "C12",
                "D1",
                "D2",
                "D3",
                "D4",
                "D5",
                "D6",
                "D7",
                "D8",
                "D9",
                "D10",
                "D11",
                "D12"
            ],
            "price": "100000",
            "description": "Khu a nè"
        },
        {
            "area": "Khu 2",
            "seats": [
                "E1",
                "E2",
                "E3",
                "E4",
                "E5",
                "E6",
                "E7",
                "E8",
                "E9",
                "E10",
                "F1",
                "F2",
                "F3",
                "F4",
                "F5",
                "F6",
                "F7",
                "F8",
                "F9",
                "F10",
                "G1",
                "G2",
                "G3",
                "G4",
                "G6",
                "G5",
                "G7",
                "G8",
                "G9",
                "G10"
            ],
            "price": "80000",
            "description": "Khu b nè"
        },
        {
            "area": "Khu 3",
            "seats": [
                "I1",
                "I2",
                "I3",
                "I4",
                "I5",
                "I6",
                "I7",
                "I8",
                "I9",
                "I10"
            ],
            "price": "50000",
            "description": "Khu c á"
        }
    ]
)

const booked = ref(
    [
        "A4",
        "A5",
        "A6",
        "B4",
        "B5",
        "B6"
    ]
)

const props = defineProps({
    data: {
        type: [],
        default: [
            {
                "area": "Khu 1",
                "seats": [
                    "A1",
                    "A2",
                    "A3"
                ],
                "price": "100000",
                "description": "Khu a nè"
            },
            {
                "area": "Khu 2",
                "seats": [
                    "A4",
                    "A5",
                    "A6"
                ],
                "price": "10000",
                "description": "Khu b nè"
            }
        ]
    },
    booked: {
        type: [],
        default: ''
    }
})

const { findNode, onConnect, setNodes, setEdges, dimensions, setTransform, nodes, edges, addEdges, addNodes, viewport, project, vueFlowRef, toObject } = useVueFlow({
})


const onDragOver = (event) => {
    event.preventDefault()
      if (event.dataTransfer) {
        event.dataTransfer.dropEffect = 'move'
      }
}

const initMap = () => {
    var vSpace = 70
    var hSpace = 70
    var x = 50
    var y = 0
    var elements = []
    for (var i = 0; i < stagemap.value.length; i++) {
        var row = stagemap.value[i]
        var label = row.label.trim()
        var number = row.number
        for (var j = 1; j <= number; j++) {
            var index = `${label}${j}`
            var areaIndex = getAreaIndex(index)
            var yPosition = y + (i * vSpace)
            var xPosition = x + (j * hSpace)
            var isBooked = checkIsBooked(index)
            var areaClassName = isBooked ? `seat-area-${areaIndex} booked` : `seat-area-${areaIndex}`
            var node = {
                id: index,
                type: 'seat',
                label: index,
                position: { x: xPosition, y: yPosition },
                class: areaClassName,
                style: { width: '42px', height: '42px' },
                events:{
                    click :() =>{
                        console.log("selected node " + index)
                    }
                }

            }
            // debugger
            // if (!isBooked) {
            //     node["events"] = {
            //         click: () => {
            //             debugger
            //             console.log("selected node " + index)
            //         }
            //     }
            // }
            elements.push(node)
        }
    }
    setNodes(elements)
}

defineExpose({
    initMap,
});


const checkIsBooked = (seat) => {
    return booked.value.includes(seat)
}


const getAreaIndex = (seat) => {
    for (var k = 0; k < exampleData.value.length; k++) {
        var seats = exampleData.value[k]["seats"]
        if (seats.includes(seat)) {
            return k + 1;
        }
    }
    return 1
}

const onSelect = (event, seat) => {
    if (event.dataTransfer) {
        event.dataTransfer.setData("seat", seat)
        event.dataTransfer.effectAllowed = 'move'
    }
}

</script>
<style lang="scss" scoped>

</style>

<style>
/* these are necessary styles for vue flow */
@import '@vue-flow/core/dist/style.css';

/* this contains the default theme, these are optional styles */
@import '@vue-flow/core/dist/theme-default.css';

.ezy-map {
    width: 70%;
    height: 440px;
}

.ezy-map-container {
    width: 100%;
    height: 440px;
}

.vue-flow__node {
    border: none !important;
    pointer-events: none !important;
}

/* body.drag * {
    pointer-events: none !important;
} */

.seat-area-1 {
    background-color: #EB4747;
}

.seat-area-2 {
    background-color: #FFB562;
}

.seat-area-3 {
    background-color: #ABC9FF;
}

.seat-area-4 {
    background-color: #B9F8D3;
}

.seat-area-5 {
    background-color: #FF5959;
}

.seat-area-6 {
    background-color: #1AA6B7;
}

.seat-area-7 {
    background-color: #6BCB77;
}

.seat-area-8 {
    background-color: #D65D7A;
}

.seat-area-9 {
    background-color: #0F3460;
}

.seat-area-10 {
    background-color: #F23557;
}

.booked {
    background-color: #EDEDED;
}
.vue-flow__node-default .vue-flow__handle, .vue-flow__node-input .vue-flow__handle, .vue-flow__node-output .vue-flow__handle {
    background: #ffffff;
}
</style>
  