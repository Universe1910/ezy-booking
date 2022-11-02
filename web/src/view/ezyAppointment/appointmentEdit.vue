
<template>
  <el-form :model="formData" ref="vForm" :rules="rules" label-position="left" label-width="150px" size="medium"
    @submit.prevent>
    <div class="static-content-item">
      <div>Add new Appointment</div>
    </div>
    <div class="static-content-item">
      <el-divider direction="horizontal"></el-divider>
    </div>
    <el-form-item label="Name" label-width="80px" prop="appointmentName" class="required">
      <el-input v-model="formData.appointmentName" size="large" type="text"
        placeholder="Please enter the appointment name" clearable></el-input>
    </el-form-item>
    <el-form-item label="Slug" label-width="80px" prop="slug">
      <el-input v-model="formData.slug" type="text" clearable><template #append>
          <el-button class="el-icon-edit"></el-button>
        </template></el-input>
    </el-form-item>
    <div class="static-content-item">
      <el-divider direction="horizontal"></el-divider>
    </div>
    <div class="static-content-item">
      <div>General</div>
    </div>
    <el-row>
      <el-col :span="18" class="grid-cell left-panel">
        <el-row>
          <el-col :span="12" class="grid-cell">
            <el-form-item label="Singer" label-width="80px" prop="singer" class="required">
              <el-input v-model="formData.singer" type="text" clearable></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12" class="grid-cell">
            <el-form-item label="Date" label-width="80px" prop="appointmentDate" class="required">
              <el-date-picker v-model="formData.appointmentDate" type="date" class="full-width-input" clearable>
              </el-date-picker>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="12" class="grid-cell">
            <el-form-item label="Start At" label-width="80px" prop="startAt" class="required">
              <el-time-picker v-model="formData.startAt" class="full-width-input" format="HH:mm:ss"
                value-format="HH:mm:ss" clearable></el-time-picker>
            </el-form-item>
          </el-col>
          <el-col :span="12" class="grid-cell">
            <el-form-item label="End At" label-width="80px" prop="endAt" class="required">
              <el-time-picker v-model="formData.endAt" class="full-width-input" format="HH:mm:ss"
                value-format="HH:mm:ss" clearable></el-time-picker>
            </el-form-item>
          </el-col>
        </el-row>
        <div class="static-content-item">
          <el-divider direction="horizontal"></el-divider>
        </div>
        <el-form-item label="Note" label-width="80px" prop="appointmentNote">
          <el-input type="textarea" v-model="formData.appointmentNote" rows="3"></el-input>
        </el-form-item>
        <el-form-item label="Content" label-width="80px" prop="appointmentContent">
          <vue-editor v-model="formData.appointmentContent"></vue-editor>
        </el-form-item>
        <div class="static-content-item">
          <el-divider direction="horizontal"></el-divider>
        </div>
        <div class="static-content-item">
          <div>Stage</div>
        </div>
        <el-form-item label="Stage" label-width="80px" prop="stage">
          <el-select v-model="formData.stage" class="full-width-input" clearable placeholder="Pick a stage">
            <el-option v-for="item in stageOptions" :key="item.ID" :label="item.name" :value="item.ID"
              :disabled="item.disabled"></el-option>
          </el-select>
        </el-form-item>
        <div class="static-content-item">
          <div style="width:100%">
            <el-table :data="tableData" style="width: 100%">
              <el-table-column label="Column">
                <template #default="scope">
                  <!-- <el-input v-model="scope.row.col" :disabled="!scope.row.edited"></el-input> -->
                  <el-input type="text" v-model="scope.row.label"></el-input>
                </template>
              </el-table-column>

              <el-table-column label="Row">
                <template #default="scope">
                  <!-- <el-input v-model="scope.row.row" :disabled="!scope.row.edited"></el-input> -->
                  <el-input type="number" @blur="changeNumberInStageMap" v-model="scope.row.number"></el-input>
                </template>
              </el-table-column>
              <el-table-column fixed="right" label="Action" width="120">
                <template #default="scope">
                  <el-button link type="primary" size="small" @click.prevent="deleteRow(scope.$index)">
                    Remove
                  </el-button>
                </template>
              </el-table-column>
            </el-table>
            <el-button class="mt-4" style="width: 100%" @click="onAddItem">Add Row</el-button>
          </div>
        </div>
        <div class="static-content-item">
          <div>Stage Area</div>
        </div>
        <div class="static-content-item" v-show="false">
          <div>Disable Index</div>
        </div>
        <el-form-item label="" label-width="0" prop="disableIndex">
          <el-select v-model="formData.disableIndex" class="full-width-input" clearable filterable allow-create
            default-first-option multiple>
            <el-option v-for="(item, index) in disableIndexOptions" :key="index" :label="item.label" :value="item.value"
              :disabled="item.disabled"></el-option>
          </el-select>
        </el-form-item>
      </el-col>
      <el-col :span="6" class="grid-cell right-panel">
        <el-form-item label="Total Seat" label-width="80px" prop="totalSeat">
          <el-input-number v-model="formData.totalSeat" class="full-width-input" :disabled="true"
            controls-position="right" :min="-100000000000" :max="100000000000" :precision="0" :step="1">
          </el-input-number>
        </el-form-item>
        <el-form-item label="Hide Index" label-width="80px" prop="hideStageIndex" class="label-right-align">
          <el-switch v-model="formData.hideStageIndex"></el-switch>
        </el-form-item>
        <el-form-item label="Allow Bus" label-width="80px" prop="allowBus">
          <el-switch v-model="formData.allowBus"></el-switch>
        </el-form-item>
        <el-form-item label="Status" label-width="80px" prop="status" class="required">
          <el-select v-model="formData.status" class="full-width-input" clearable>
            <el-option v-for="(item, index) in statusOptions" :key="index" :label="item.label" :value="item.value"
              :disabled="item.disabled"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="Author" label-width="80px" prop="createdBy" class="required">
          <el-select v-model="formData.createdBy" class="full-width-input" clearable filterable automatic-dropdown
            remote>
            <el-option v-for="(item, index) in authorOptions" :key="index" :label="item.label" :value="item.value"
              :disabled="item.disabled"></el-option>
          </el-select>
          <el-form-item>
            <el-button @click="updateAppointment" type="primary">Update</el-button>
          </el-form-item>

        </el-form-item>
        <div class="static-content-item">
          <el-divider direction="horizontal"></el-divider>
        </div>
        <div class="static-content-item">
          <div>Image</div>
        </div>
        <el-form-item label="" label-width="0" prop="featured_image" class="label-right-align">
          <el-upload :file-list="featured_imageFileList" :headers="featured_imageUploadHeaders"
            :data="featured_imageUploadData" list-type="picture-card" with-credentials show-file-list :limit="1">
            <template #default><i class="el-icon-plus"></i></template>
          </el-upload>
        </el-form-item>
        <div class="static-content-item">
          <el-divider direction="horizontal"></el-divider>
        </div>
        <el-form-item label="Branch" label-width="80px" prop="branch">
          <el-radio-group v-model="formData.branch">
            <el-radio v-for="item in branchOptions" :key="item.ID" :label="item.name" :disabled="item.disabled"
              style="{display: inline}">{{ item.name }}</el-radio>
          </el-radio-group>
        </el-form-item>
      </el-col>
    </el-row>
  </el-form>

</template>

<script>
export default {
  name: 'EzyAppointmentEdit'
}
</script>

<script setup>
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { useRoute } from 'vue-router'

import {
  updateEzyAppointment,
  findEzyAppointment,
} from '@/api/ezyAppointment'

import {
  getEzyBranchList
} from '@/api/ezyBranch'

import {
  getEzyStageList
} from '@/api/ezyStage'

import {
  getUserList,
} from '@/api/user'

const route = useRoute()

const formData = ref({
  appointmentName: "",
  slug: "",
  singer: "",
  appointmentDate: null,
  startAt: null,
  endAt: null,
  appointmentNote: "",
  appointmentContent: null,
  stage: "",
  stageMap: "",
  stageArea: "",
  disableIndex: [],
  totalSeat: 0,
  hideStageIndex: false,
  allowBus: false,
  status: "publish",
  createdBy: "",
  featuredImage: null,
  branch: 3,

},
)



const rule = reactive(
  {
    appointmentName: [{
      required: true,
      message: 'Input value should be not null.',
    }],
    slug: [{
      pattern: /^([hH][tT]{2}[pP]:\/\/|[hH][tT]{2}[pP][sS]:\/\/)(([A-Za-z0-9-~]+)\.)+([A-Za-z0-9-~\/])+$/,
      trigger: ['blur', 'change'],
      message: ''
    }],
    singer: [{
      required: true,
      message: 'Input value should be not null.',
    }],
    appointmentDate: [{
      required: true,
      message: 'Input value should be not null.',
    }],
    startAt: [{
      required: true,
      message: 'Input value should be not null.',
    }],
    endAt: [{
      required: true,
      message: 'Input value should be not null.',
    }],
    status: [{
      required: true,
      message: 'Input value should be not null.',
    }],
    createdBy: [{
      required: true,
      message: 'Input value should be not null.',
    }],
  },
)

const statusOptions = reactive(
  [
    {
      label: 'Publish',
      value: 'publish'
    },
    {
      label: 'Draft',
      value: 'draft'
    },
    {
      label: 'End',
      value: 'end'
    },
    {
      label: 'View',
      value: 'view'
    },
    {
      label: 'Request Admin',
      value: 'request_admin'
    },
    {
      label: 'Trash',
      value: 'trash'
    }
  ]
)

const stageOptions = ref([])

const authorOptions = ref([])

const branchOptions = ref([])

const featured_imageFileList = ref([]);

const featured_imageUploadHeaders = ref({});

const featured_imageUploadData = ref({});

const disableIndexOptions = ref([]);

const stageMapOptions = ref([]);


const tableData = ref([])

const searchInfo = ref({ appointmentId: Number(route.params.id) })


const getEzyAppointmentById = async () => {
  let appointmentId = searchInfo.value.appointmentId;
  var res = await findEzyAppointment({ ID: appointmentId });
  console.log(res);
  formData.value = res.data.reezyAppointment;
}

getEzyAppointmentById();


// const getAuthors = async () => {
//   const table = await getUserList()
//   if (table.code === 0) {
//     authorOptions.value = table.data.list
//     console.log("authors")
//     console.log(authorOptions.value)
//   }
// }

// getAuthors()

const getBranchs = async () => {
  const table = await getEzyBranchList()
  if (table.code === 0) {
    branchOptions.value = table.data.list
    console.log("branchs")
    console.log(branchOptions.value)
  }
}
getBranchs()

const getStages = async () => {
  const table = await getEzyStageList()
  if (table.code === 0) {
    stageOptions.value = table.data.list
    console.log("stages")
    console.log(stageOptions.value)
  }
}

getStages()

const deleteRow = (index) => {
  tableData.value.splice(index, 1)
  changeNumberInStageMap();
}

const onAddItem = () => {
  tableData.value.push({
    stage_map_col: '',
    stage_map_row: 0,
  })
  // console.log(tableData.value);
  calcTotalSeat();
}

const changeNumberInStageMap = () => {
  var total = 0;
  debugger;
  for (var i = 0; i < tableData.value.length; i++) {
    var row = tableData.value[i];
    var number = row['number'];
    if (number) {
      total += number * 1;
    }
  }
  //  tableData.value.forEach((e) =>{
  //       total += e.number * 1;
  // });
  console.log("total: " + total);
  formData.value.totalSeat = total

}

const handleEdit = (index, row) => {
  console.log(index, row)
}

//  const instance = getCurrentInstance()

const submitForm = () => {
  // instance.ctx.$refs['vForm'].validate(valid => {
  //   if (!valid) return
  //   //TODO: 提交表单
  // })

  const resetForm = () => {
    // instance.ctx.$refs['vForm'].resetFields()
  }


}

</script>

<style lang="scss">
.el-input-number.full-width-input,
.el-cascader.full-width-input {
  width: 100% !important;
}

.el-form-item--medium {
  .el-radio {
    line-height: 36px !important;
  }

  .el-rate {
    margin-top: 8px;
  }
}

.el-form-item--small {
  .el-radio {
    line-height: 32px !important;
  }

  .el-rate {
    margin-top: 6px;
  }
}

.el-form-item--mini {
  .el-radio {
    line-height: 28px !important;
  }

  .el-rate {
    margin-top: 4px;
  }
}

.clear-fix:before,
.clear-fix:after {
  display: table;
  content: "";
}

.clear-fix:after {
  clear: both;
}

.float-right {
  float: right;
}
</style>

<style lang="scss" scoped>
div.table-container {
  table.table-layout {
    width: 100%;
    table-layout: fixed;
    border-collapse: collapse;

    td.table-cell {
      display: table-cell;
      height: 36px;
      border: 1px solid #e1e2e3;
    }
  }
}

div.tab-container {}

.label-left-align :deep(.el-form-item__label) {
  text-align: left;
  padding: 8px;
}

.label-center-align :deep(.el-form-item__label) {
  text-align: center;
}

.label-right-align :deep(.el-form-item__label) {
  text-align: right;
}

.custom-label {}

.static-content-item {
  min-height: 20px;
  display: flex;
  align-items: center;

  :deep(.el-divider--horizontal) {
    margin: 0;
  }
}

.left-panel {
  padding-right: 6px;
  border-right: solid 1px #DDDFE6;
}

.right-panel {
  padding-left: 6px;

}
</style>