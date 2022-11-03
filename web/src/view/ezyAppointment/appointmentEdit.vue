
<template>
  <el-form :model="formData" ref="vForm" :rules="rule" label-position="left" label-width="150px" size="medium">
    <div class="static-content-item">
      <h3>Add new Appointment</h3>
    </div>
    <div class="static-content-item">
      <el-divider direction="horizontal"></el-divider>
    </div>
    <el-form-item label="Name" label-width="80px" prop="appointmentName" class="required">
      <el-input v-model="formData.appointmentName" size="large" type="text"
        placeholder="Please enter the appointment name" clearable></el-input>
    </el-form-item>
    <el-form-item label="Slug" label-width="80px" prop="slug">
      <el-input v-model="formData.slug" type="text" clearable :disabled="true"><template #append>
          <el-button class="el-icon-edit"></el-button>
        </template></el-input>
    </el-form-item>
    <div class="static-content-item">
      <el-divider direction="horizontal"></el-divider>
    </div>
    <div class="static-content-item">
      <h3>General</h3>
    </div>
    <el-row>
      <el-col :span="18" class="grid-cell left-panel">
        <el-row>
          <el-col :span="12" class="grid-cell">
            <el-form-item label="Singer" label-width="80px" prop="singer" class="required singer-input">
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
        <el-form-item label="Content" label-width="80px" prop="appointmentContent">
          <vue-editor class="appointment-content-editor" v-model="formData.appointmentContent"></vue-editor>
        </el-form-item>
        <el-form-item label="Note" label-width="80px" prop="appointmentNote">
          <el-input type="textarea" v-model="formData.appointmentNote" rows="5"></el-input>
        </el-form-item>

        <div class="static-content-item">
          <el-divider direction="horizontal">
          </el-divider>
        </div>
        <div class="static-content-item">
          <h3>Stage</h3>
        </div>
        <div class="static-content-item">
        </div>
        <el-form-item label="Stage" label-width="80px" prop="stageId">
          <el-select v-model="formData.stageId" class="full-width-input" clearable placeholder="Pick a stage">
            <el-option v-for="item in stageOptions" :key="item.ID" :label="item.name" :value="item.ID"
              :disabled="item.disabled"></el-option>
          </el-select>
        </el-form-item>
        <div class="static-content-item">
          <div style="width:100%">
            <el-table :data="tableData" style="width: 100%">
              <el-table-column label="Column">
                <template #default="scope">
                  <el-input type="text" v-model="scope.row.label"></el-input>
                </template>
              </el-table-column>
              <el-table-column label="Row">
                <template #default="scope">
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

        <el-form-item label="Disable" label-width="80px" prop="disableIndex">
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
        <el-form-item label="Branch" label-width="80px" prop="branchId">
          <el-select v-model="formData.branchId" class="full-width-input" clearable placeholder="Pick a branch">
            <el-option v-for="item in branchOptions" :key="item.ID" :label="item.name" :value="item.ID"
              :disabled="item.disabled"></el-option>
          </el-select>
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
            <el-option v-for="item in authorOptions" :key="item.ID" :label="item.nickName" :value="item.ID"
              :disabled="item.disabled"></el-option>
          </el-select>
          <el-form-item class="update-button">
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

import { VueEditor } from "vue3-editor";


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

const vForm = ref()
const formData = ref({
  appointmentName: "",
  slug: "",
  singer: "",
  appointmentDate: null,
  startAt: null,
  endAt: null,
  appointmentNote: "",
  appointmentContent: null,
  stageId: "",
  branchId: "",
  stageMap: "",
  stageArea: "",
  disableIndex: [],
  totalSeat: 0,
  hideStageIndex: false,
  allowBus: false,
  status: "publish",
  createdBy: 1,
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
    // slug: [{
    //   pattern: /^([hH][tT]{2}[pP]:\/\/|[hH][tT]{2}[pP][sS]:\/\/)(([A-Za-z0-9-~]+)\.)+([A-Za-z0-9-~\/])+$/,
    //   trigger: ['blur', 'change'],
    //   message: ''
    // }],
    stageId: [{
      required: true,
      message: 'Input value should be not null.',
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

  formData.value = res.data.reezyAppointment;
  debugger;
  if (formData.value.stageMap) {
    var stageMapObject = JSON.parse(formData.value.stageMap)
    tableData.value = stageMapObject
  }
  console.log(formData.value);
}


getEzyAppointmentById();


const getAuthors = async () => {
  const table = await getUserList({ page: 1, pageSize: 100 })
  if (table.code === 0) {
    authorOptions.value = table.data.list
    // console.log("authors")
    console.log(authorOptions.value)
  }
}

getAuthors()

const getBranchs = async () => {
  const table = await getEzyBranchList()
  if (table.code === 0) {
    branchOptions.value = table.data.list
    // console.log("branchs")
    // console.log(branchOptions.value)
  }
}
getBranchs()

const getStages = async () => {
  const table = await getEzyStageList()
  if (table.code === 0) {
    stageOptions.value = table.data.list
  }
}

getStages()

const updateAppointment = async () => {
  vForm.value?.validate(async (valid) => {
    console.log("update appointment")
    if (!valid) return
    debugger;
    var slug = toSlug(formData.value.appointmentName)
    formData.value.slug = slug;
    var stageMapObject = getStageMap()
    formData.value.stageMap = JSON.stringify(stageMapObject)
    var res = await updateEzyAppointment(formData.value)
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: 'Update successfullys'
      })
      // closeDialog()
      // getTableData()
    }
  })
}

const getStageMap = () => {
  debugger;
  var stageMapObject = []
  for (var i = 0; i < tableData.value.length; i++) {
    var row = tableData.value[i];
    stageMapObject.push({
      label: row['label'],
      number: row['number'],
    })
  }
  return stageMapObject
}

const deleteRow = (index) => {
  tableData.value.splice(index, 1)
  changeNumberInStageMap();
}

const onAddItem = () => {
  tableData.value.push({
    lable: '',
    row: 0,
  })
  calcTotalSeat();
}

const changeNumberInStageMap = () => {
  var total = 0;
  for (var i = 0; i < tableData.value.length; i++) {
    var row = tableData.value[i];
    var number = row['number'];
    if (number) {
      total += number * 1;
    }
  }
  formData.value.totalSeat = total

}

const handleEdit = (index, row) => {
  console.log(index, row)
}


const toSlug = (str) => {
  // Chuyển hết sang chữ thường
  str = str.toLowerCase();

  // xóa dấu
  str = str
    .normalize('NFD') // chuyển chuỗi sang unicode tổ hợp
    .replace(/[\u0300-\u036f]/g, ''); // xóa các ký tự dấu sau khi tách tổ hợp

  // Thay ký tự đĐ
  str = str.replace(/[đĐ]/g, 'd');

  // Xóa ký tự đặc biệt
  str = str.replace(/([^0-9a-z-\s])/g, '');

  // Xóa khoảng trắng thay bằng ký tự -
  str = str.replace(/(\s+)/g, '-');

  // Xóa ký tự - liên tiếp
  str = str.replace(/-+/g, '-');

  // xóa phần dư - ở đầu & cuối
  str = str.replace(/^-+|-+$/g, '');

  // return
  return str;
}
</script>


<style lang="scss">

</style>

<style lang="scss" scoped>
.el-input-number.full-width-input,
.el-cascader.full-width-input {
  width: 100% !important;
}

.el-select {
  width: 100% !important;
}

.el-form-item--medium {
  width: 100%;

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
  // margin-top: 12px;
  display: flex;
  align-items: center;

  :deep(.el-divider--horizontal) {
    margin: 4;
  }
}

.appointment-content-editor {
  background-color: #fff;
  border-radius: 8px;
}

.left-panel {
  padding-right: 6px;
  border-right: solid 1px #DDDFE6;
}

.right-panel {
  padding-left: 6px;
}

.update-button {
  margin-top: 16px;
}

.singer-input {
  width: 80%;
}
</style>