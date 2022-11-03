<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="Appointment Name:" prop="appointmentName">
          <el-input v-model="formData.appointmentName" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="Singer:" prop="singer">
          <el-input v-model="formData.singer" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="Appointment Date:" prop="appointmentDate">
          <el-date-picker v-model="formData.appointmentDate" type="date" placeholder="选择日期" :clearable="false"></el-date-picker>
        </el-form-item>
        <el-form-item label="Start At:" prop="startAt">
          <el-date-picker v-model="formData.startAt" type="date" placeholder="选择日期" :clearable="false"></el-date-picker>
        </el-form-item>
        <el-form-item label="End At:" prop="endAt">
          <el-date-picker v-model="formData.endAt" type="date" placeholder="选择日期" :clearable="false"></el-date-picker>
        </el-form-item>
        <el-form-item label="Appointment Content:" prop="appointmentContent">
          <el-input v-model="formData.appointmentContent" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="Appointment Note:" prop="appointmentNote">
          <el-input v-model="formData.appointmentNote" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="Stage:" prop="stage">
          <el-input v-model.number="formData.stage" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="Stage Map:" prop="stageMap">
          <el-input v-model="formData.stageMap" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="Stage Area:" prop="stageArea">
          <el-input v-model="formData.stageArea" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="Disable Index:" prop="disableIndex">
          <el-input v-model="formData.disableIndex" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="Total Seat:" prop="totalSeat">
          <el-input v-model.number="formData.totalSeat" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="Hide Index:" prop="hideStageIndex">
          <el-switch v-model="formData.hideStageIndex" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
        </el-form-item>
        <el-form-item label="Allow Bus:" prop="allowBus">
          <el-switch v-model="formData.allowBus" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
        </el-form-item>
        <el-form-item label="CreatedBy:" prop="createdBy">
          <el-input v-model.number="formData.createdBy" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="Branch:" prop="branch">
          <el-input v-model.number="formData.branch" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="Featured Image:" prop="featuredImage">
          <el-input v-model="formData.featuredImage" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="Status:" prop="status">
          <el-input v-model="formData.status" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" @click="save">保存</el-button>
          <el-button size="small" type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'EzyAppointment'
}
</script>

<script setup>
import {
  createEzyAppointment,
  updateEzyAppointment,
  findEzyAppointment
} from '@/api/ezyAppointment'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            appointmentName: '',
            singer: '',
            appointmentDate: new Date(),
            startAt: new Date(),
            endAt: new Date(),
            appointmentContent: '',
            appointmentNote: '',
            stage: 0,
            stageMap: '',
            stageArea: '',
            disableIndex: '',
            totalSeat: 0,
            hideStageIndex: false,
            allowBus: false,
            createdBy: 0,
            branch: 0,
            featuredImage: '',
            status: '',
        })
// 验证规则
const rule = reactive({
               appointmentName : [{
                   required: true,
                   message: 'Appointment name can\'t be empty',
                   trigger: ['input','blur'],
               }],
               singer : [{
                   required: true,
                   message: 'Singer can\'t be empty',
                   trigger: ['input','blur'],
               }],
               appointmentDate : [{
                   required: true,
                   message: 'Appointment date can\'t be empty',
                   trigger: ['input','blur'],
               }],
               startAt : [{
                   required: true,
                   message: 'Start at can\'t be empty',
                   trigger: ['input','blur'],
               }],
               endAt : [{
                   required: true,
                   message: 'End at can\'t be empty',
                   trigger: ['input','blur'],
               }],
               stage : [{
                   required: true,
                   message: 'Stage can\'t be empty',
                   trigger: ['input','blur'],
               }],
               stageMap : [{
                   required: true,
                   message: 'Stage map can\'t be empty',
                   trigger: ['input','blur'],
               }],
               stageArea : [{
                   required: true,
                   message: 'Stage Area can\'t be empty',
                   trigger: ['input','blur'],
               }],
               branch : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               status : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findEzyAppointment({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reezyAppointment
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createEzyAppointment(formData.value)
               break
             case 'update':
               res = await updateEzyAppointment(formData.value)
               break
             default:
               res = await createEzyAppointment(formData.value)
               break
           }
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: 'Create /更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
