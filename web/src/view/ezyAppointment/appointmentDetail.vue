<template>
  <el-form :model="appointmentData" ref="formRef" :rules="rules" label-position="top" label-width="120px" size="medium"
    @submit.prevent>
    <el-row>
      <el-col :span="12" class="grid-cell">
        <div class="static-content-item">
          <div>Concert Mưa</div>
        </div>
        <div class="static-content-item">
          <el-divider direction="horizontal"></el-divider>
        </div>
        <div class="static-content-item">
          <div>Published</div>
        </div>
      </el-col>
      <el-col :span="12" class="grid-cell">
        <div class="static-content-item">
          <div>Singer</div>
        </div>
        <div class="static-content-item">
          <div>Hà Anh Tuấn</div>
        </div>
        <div class="static-content-item">
          <div>Date</div>
        </div>
        <div class="static-content-item">
          <div>31/12/2022</div>
        </div>
        <div class="static-content-item">
          <div>Time</div>
        </div>
        <div class="static-content-item">
          <div>17:00</div>
        </div>
      </el-col>
    </el-row>
    <div class="static-content-item">
      <el-divider direction="horizontal"></el-divider>
    </div>
    <el-row>
      <el-col :span="3" class="grid-cell">
        <div class="static-content-item">
          <el-button type="info">View</el-button>
        </div>
      </el-col>
      <el-col :span="3" class="grid-cell">
        <div class="static-content-item">
          <el-button type="primary">Booking</el-button>
        </div>
      </el-col>
      <el-col :span="3" class="grid-cell">
        <div class="static-content-item">
          <el-button type="primary" icon="download">Excel</el-button>
        </div>
      </el-col>
      <el-col :span="3" class="grid-cell">
        <div class="static-content-item">
          <el-button type="primary">Refresh</el-button>
        </div>
      </el-col>
      <el-col :span="3" class="grid-cell">
        <div class="static-content-item">
          <el-button type="primary">Edit</el-button>
        </div>
      </el-col>
      <el-col :span="3" class="grid-cell">
        <div class="static-content-item">
          <el-button type="success">Stats</el-button>
        </div>
      </el-col>
      <el-col :span="3" class="grid-cell">
        <div class="static-content-item">
          <el-button type="danger">End</el-button>
        </div>
      </el-col>
      <el-col :span="3" class="grid-cell">
        <div class="static-content-item" v-show="false">
          <el-button :disabled="true">Advance</el-button>
        </div>
      </el-col>
    </el-row>

    <el-row class="h800">
      <el-col :span="16" class="grid-cell">
        <EzyMap ref="ezyMapRef" :data="stageAreaData" :booked="booked"></EzyMap>
      </el-col>
      <el-col :span="8" class="grid-cell padding-8">

        <div class="booking-form">
          <el-form :model="bookingForm" ref="vBookingForm" :rules="bookingFormRules" label-position="top" size="medium">
            <el-form-item label="Customer Name" prop="customerName" class="required">
              <el-input v-model="bookingForm.customerName" size="large" type="text"
                placeholder="Please enter the customer name" clearable></el-input>
            </el-form-item>
            <el-form-item label="Customer Phone" prop="customerPhone" class="required">
              <el-input v-model="bookingForm.customerPhone" size="large" type="text" placeholder="Phone number"
                clearable></el-input>
            </el-form-item>

            <el-form-item label="Customer Email" prop="customerEmail" class="required">
              <el-input v-model="bookingForm.customerEmail" size="large" type="email" placeholder="Email"
                clearable></el-input>
            </el-form-item>
            <el-form-item label="Seats" prop="seats" class="required w100">
              <!-- <el-select v-model="bookingForm.seats" size="large" multiple clearable></el-select> -->
            </el-form-item>
            <div class="form-item-row">
              <el-form-item label="Coupon" prop="couponCode">
                <el-select v-model="bookingForm.couponCode" size="large" clearable>
                </el-select>
              </el-form-item>
              <el-form-item label="Membership Code" prop="membershipCode" class="membership-code-input">
                <el-input v-model="bookingForm.membershipCode" size="large" type="text" clearable></el-input>
              </el-form-item>
            </div>

            <div class="form-item-row">
              <el-form-item label="Admin Discount" prop="adminDiscount">
                <el-input @change="calculateInvoicePayment" type="number" v-model="bookingForm.adminDiscount"
                  size="large" clearable>
                </el-input>
              </el-form-item>
              <el-form-item label="Tax" prop="tax">
                <el-input type="number" @change="calculateInvoicePayment" v-model="bookingForm.tax" size="large"
                  clearable>
                </el-input>
              </el-form-item>
            </div>

            <el-form-item label="Payment Type" prop="paymentType" class="required">
              <el-select v-model="bookingForm.paymentType" size="large" clearable>
                <el-option value="Direct bank transfer">Direct bank transfer</el-option>
                <el-option value="POS">POS</el-option>
                <el-option value="Online bank transfer">Online bank transfer</el-option>
              </el-select>
            </el-form-item>
            <el-divider></el-divider>
            <div class="booking-sumary">
              <el-row>
                <el-col :span="8" class="grid-cell">Total</el-col>
                <el-col :span="12" class="grid-cell">
                  <div class="booking-total">{{ formatCurrency(bookingForm.total * 1) }}</div>
                </el-col>
              </el-row>
              <el-row v-if="bookingForm.adminDiscount">
                <el-col :span="8" class="grid-cell">Admin Discount</el-col>
                <el-col :span="12" class="grid-cell">
                  <div class="booking-admin-discount">- {{ formatCurrency(bookingForm.adminDiscount * 1) }}</div>
                </el-col>
              </el-row>
              <el-row v-if="bookingForm.couponDiscount">
                <el-col :span="8" class="grid-cell">Coupon Discount</el-col>
                <el-col :span="12" class="grid-cell">
                  <div class="booking-coupon-discount">- {{ formatCurrency(bookingForm.couponDiscount * 1) }}</div>
                </el-col>
              </el-row>
              <el-row v-if="bookingForm.membershipDiscount">
                <el-col :span="8" class="grid-cell">Membership Discount</el-col>
                <el-col :span="12" class="grid-cell">
                  <div class="booking-membership-discount">- {{ formatCurrency(bookingForm.membershipDiscount * 1) }}
                  </div>
                </el-col>
              </el-row>
              <el-row v-if="bookingForm.tax">
                <el-col :span="8" class="grid-cell">Tax</el-col>
                <el-col :span="12" class="grid-cell">
                  <div class="booking-tax">+ {{ formatCurrency(bookingForm.tax * 1) }}</div>
                </el-col>
              </el-row>
              <el-divider></el-divider>
              <el-row>
                <el-col :span="8" class="grid-cell">Invoice Payment</el-col>
                <el-col :span="12" class="grid-cell">
                  <div class="booking-invoice-payment">{{ formatCurrency(bookingForm.invoicePayment * 1) }}</div>
                </el-col>
              </el-row>
            </div>

          </el-form>
          <el-button type="primary" size="langer" class="submit-butotn" @click="createOrder">Order</el-button>
        </div>
      </el-col>
    </el-row>

    <div class="static-content-item">
      <el-divider direction="horizontal"></el-divider>
    </div>
    <el-row>
      <el-col :span="4" class="grid-cell">
        <div class="static-content-item">
          <div>Total seat</div>
        </div>
        <div class="static-content-item">
          <div>650 ghế</div>
        </div>
      </el-col>
      <el-col :span="4" class="grid-cell">
        <div class="static-content-item">
          <div>Lock</div>
        </div>
        <div class="static-content-item">
          <div>228 ghế</div>
        </div>
      </el-col>
      <el-col :span="4" class="grid-cell">
        <div class="static-content-item">
          <div>Paid</div>
        </div>
        <div class="static-content-item">
          <div>13 ghế</div>
        </div>
      </el-col>
      <el-col :span="4" class="grid-cell">
        <div class="static-content-item">
          <div>Booking</div>
        </div>
        <div class="static-content-item">
          <div>0 ghế</div>
        </div>
      </el-col>
      <el-col :span="4" class="grid-cell">
        <div class="static-content-item">
          <div>Avaliable</div>
        </div>
        <div class="static-content-item">
          <div>409 ghế</div>
        </div>
      </el-col>
      <el-col :span="4" class="grid-cell">
        <div class="static-content-item">
          <div>Turnover</div>
        </div>
        <div class="static-content-item">
          <div>13,995,000 VNĐ</div>
        </div>
      </el-col>
    </el-row>
    <el-row>
      <el-col :span="4" class="grid-cell">
        <div class="static-content-item">
          <div>Bus</div>
        </div>
        <div class="static-content-item">
          <div>13</div>
        </div>
      </el-col>
      <el-col :span="4" class="grid-cell">
        <div class="static-content-item">
          <div>Arrived</div>
        </div>
        <div class="static-content-item">
          <div>28</div>
        </div>
      </el-col>
      <el-col :span="4" class="grid-cell">
      </el-col>
      <el-col :span="4" class="grid-cell">
      </el-col>
      <el-col :span="4" class="grid-cell">
      </el-col>
      <el-col :span="4" class="grid-cell">
      </el-col>
    </el-row>
    <div class="static-content-item">
      <el-divider direction="horizontal"></el-divider>
    </div>
    <div class="table-container">
      <table class="table-layout">
        <tbody>
          <tr>
            <td class="table-cell">
              <div class="static-content-item">
                <div>AREA</div>
              </div>
            </td>
            <td class="table-cell">
              <div class="static-content-item">
                <div>NUMBER</div>
              </div>
            </td>
            <td class="table-cell">
              <div class="static-content-item">
                <div>PRICE</div>
              </div>
            </td>
            <td class="table-cell">
              <div class="static-content-item">
                <div>TOTAL</div>
              </div>
            </td>
            <td class="table-cell">
              <div class="static-content-item">
                <div>PAID</div>
              </div>
            </td>
            <td class="table-cell">
              <div class="static-content-item">
                <div>INTO</div>
              </div>
            </td>
          </tr>
          <tr>
            <td class="table-cell">
            </td>
            <td class="table-cell">
            </td>
            <td class="table-cell">
            </td>
            <td class="table-cell">
            </td>
            <td class="table-cell">
            </td>
            <td class="table-cell">
            </td>
          </tr>
        </tbody>
      </table>
    </div>
    <div class="static-content-item">
      <el-divider direction="horizontal"></el-divider>
    </div>
    <div class="tab-container">
      <el-tabs v-model="activeTab" @tab-click="handleTabClick" type="border-card">
        <el-tab-pane name="Completed" label="Completed">
          <el-table ref="multipleTable" style="width: 100%" tooltip-effect="dark" :data="tableData" row-key="ID"
            @selection-change="handleSelectionChange">
            <el-table-column align="left" label="ID" prop="ID" width="60" />
            <el-table-column align="left" label="Name" prop="customerObject.name" width="200" />
            <el-table-column align="left" label="Phone" prop="customerObject.phone" width="150" />
            <el-table-column align="left" label="Qty" prop="quantity" width="50" />
            <el-table-column align="left" label="Bus" prop="busNumber" width="50" />
            <el-table-column align="left" label="Arrived" prop="arrivedNumber" width="80" />
            <el-table-column align="left" label="Seats" prop="seats" width="120" />
            <el-table-column align="left" label="Total" prop="total" width="120">
              <template #default="scope">{{ formatCurrency(scope.row.total) }}</template>
            </el-table-column>
            <el-table-column align="left" label="Coupon Code" prop="couponCode" width="120" />
            <el-table-column align="left" label="Discount" prop="couponDiscount" width="120" />
            <el-table-column align="left" label="Tax" prop="tax" width="120">
              <template #default="scope">{{ formatCurrency(scope.row.tax) }}</template>
            </el-table-column>
            <el-table-column align="left" label="Payment" prop="invoicePayment" width="120">
              <template #default="scope">{{ formatCurrency(scope.row.invoicePayment) }}</template>
            </el-table-column>
            <!-- <el-table-column align="left" label="Source" prop="source" width="120" /> -->
            <el-table-column align="left" label="Notes" prop="notes" width="120" />
            <!-- <el-table-column align="left" label="Status" prop="status" width="120" /> -->
            <el-table-column align="left" label="Last Action" prop="lastActionByObject.nickName" width="120" />
            <el-table-column align="left" label="Last Edit" width="180">
              <template #default="scope">{{ formatDate(scope.row.lastActionTime) }}</template>
            </el-table-column>
            <el-table-column align="left" label="Created By" prop="createdByObject.nickName" width="120">
              <template #default="scope">{{
                  scope.row.createdBy == null ? 'Customer' :
                    scope.row.createdByObject.nickName
              }}
              </template>
            </el-table-column>
            <el-table-column align="left" label="Booking Time" width="180">
              <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
            </el-table-column>
            <el-table-column align="left" label="Action" width="120">
              <template #default="scope">
                <el-button type="danger" size="small" class="table-button" @click="cancelFromCompleted(scope.row)">
                  Cancel
                </el-button>
                <el-button type="primary" size="small" class="table-button" @click="refundlFromCompleted(scope.row)">
                  Refund
                </el-button>
                <el-button type="warning" size="small" class="table-button" @click="reserveFromCompleted(scope.row)">
                  Reserve
                </el-button>
              </template>
            </el-table-column>
          </el-table>
          <div class="gva-pagination">
            <el-pagination layout="total, sizes, prev, pager, next, jumper" :current-page="page" :page-size="pageSize"
              :page-sizes="[10, 30, 50, 100]" :total="total" @current-change="handleCurrentChange"
              @size-change="handleSizeChange" />
          </div>
        </el-tab-pane>
        <el-tab-pane name="Admin" label="Admin Booking">
          <el-table ref="multipleTable" style="width: 100%" tooltip-effect="dark" :data="tableData" row-key="ID"
            @selection-change="handleSelectionChange">
            <el-table-column align="left" label="ID" prop="ID" width="60" />
            <el-table-column align="left" label="Name" prop="customerObject.name" width="200" />
            <el-table-column align="left" label="Phone" prop="customerObject.phone" width="150" />
            <el-table-column align="left" label="Qty" prop="quantity" width="50" />
            <el-table-column align="left" label="Bus" prop="busNumber" width="50" />
            <el-table-column align="left" label="Arrived" prop="arrivedNumber" width="80" />
            <el-table-column align="left" label="Seats" prop="seats" width="120" />
            <el-table-column align="left" label="Total" prop="total" width="120">
              <template #default="scope">{{ formatCurrency(scope.row.total) }}</template>
            </el-table-column>
            <el-table-column align="left" label="Coupon Code" prop="couponCode" width="120" />
            <el-table-column align="left" label="Discount" prop="couponDiscount" width="120" />
            <el-table-column align="left" label="Tax" prop="tax" width="120">
              <template #default="scope">{{ formatCurrency(scope.row.tax) }}</template>
            </el-table-column>
            <el-table-column align="left" label="Payment" prop="invoicePayment" width="120">
              <template #default="scope">{{ formatCurrency(scope.row.invoicePayment) }}</template>
            </el-table-column>
            <!-- <el-table-column align="left" label="Source" prop="source" width="120" /> -->
            <el-table-column align="left" label="Notes" prop="notes" width="120" />
            <!-- <el-table-column align="left" label="Status" prop="status" width="120" /> -->
            <el-table-column align="left" label="Last Action" prop="lastActionByObject.nickName" width="120" />
            <el-table-column align="left" label="Last Edit" width="180">
              <template #default="scope">{{ formatDate(scope.row.lastActionTime) }}</template>
            </el-table-column>
            <el-table-column align="left" label="Created By" prop="createdByObject.nickName" width="120">
              <template #default="scope">{{
                  scope.row.createdBy == null ? 'Customer' :
                    scope.row.createdByObject.nickName
              }}
              </template>
            </el-table-column>
            <el-table-column align="left" label="Booking Time" width="180">
              <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
            </el-table-column>
            <el-table-column align="left" label="Action" width="120">
              <template #default="scope">
                <el-button type="primary" link size="small" class="table-button"
                  @click="confirmAdminBookingToCompleted(scope.row)">Confirm
                </el-button>
                <el-button type="danger" size="small" class="table-button" @click="cancelFromCompleted(scope.row)">
                  Cancel
                </el-button>

              </template>
            </el-table-column>
          </el-table>
          <div class="gva-pagination">
            <el-pagination layout="total, sizes, prev, pager, next, jumper" :current-page="page" :page-size="pageSize"
              :page-sizes="[10, 30, 50, 100]" :total="total" @current-change="handleCurrentChange"
              @size-change="handleSizeChange" />
          </div>
        </el-tab-pane>
        <el-tab-pane name="Cancel" label="Cancel">
          <el-table ref="multipleTable" style="width: 100%" tooltip-effect="dark" :data="tableData" row-key="ID"
            @selection-change="handleSelectionChange">
            <!-- <el-table-column type="selection" width="55" /> -->
            <el-table-column align="left" label="ID" prop="ID" width="60" />
            <el-table-column align="left" label="Name" prop="customerObject.name" width="200" />
            <el-table-column align="left" label="Phone" prop="customerObject.phone" width="150" />
            <!-- <el-table-column align="left" label="Email" prop="customerObject.email" width="180" /> -->
            <el-table-column align="left" label="Qty" prop="quantity" width="50" />
            <el-table-column align="left" label="Bus" prop="busNumber" width="50" />
            <el-table-column align="left" label="Seats" prop="seats" width="120" />
            <el-table-column align="left" label="Total" prop="total" width="120">
              <template #default="scope">{{ formatCurrency(scope.row.total) }}</template>
            </el-table-column>
            <el-table-column align="left" label="Coupon Code" prop="couponCode" width="120" />
            <el-table-column align="left" label="Discount" prop="couponDiscount" width="120" />
            <el-table-column align="left" label="Tax" prop="tax" width="120">
              <template #default="scope">{{ formatCurrency(scope.row.tax) }}</template>
            </el-table-column>
            <el-table-column align="left" label="Payment" prop="invoicePayment" width="120">
              <template #default="scope">{{ formatCurrency(scope.row.invoicePayment) }}</template>
            </el-table-column>
            <el-table-column align="left" label="Notes" prop="notes" width="120" />
            <!-- <el-table-column align="left" label="Status" prop="status" width="120" /> -->
            <el-table-column align="left" label="Last Action" prop="lastActionByObject.nickName" width="120" />
            <el-table-column align="left" label="Last Edit" width="180">
              <template #default="scope">{{ formatDate(scope.row.lastActionTime) }}</template>
            </el-table-column>
            <el-table-column align="left" label="Created By" prop="createdByObject.nickName" width="120">
              <template #default="scope">{{
                  scope.row.createdBy == null ? 'Customer' :
                    scope.row.createdByObject.nickName
              }}
              </template>
            </el-table-column>
            <el-table-column align="left" label="Booking Time" width="180">
              <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
            </el-table-column>
            <el-table-column align="left" label="Action">
              <template #default="scope">
                <el-button type="primary" link size="small" class="table-button"
                  @click="confirmProcessingToCompleted(scope.row)">Confirm
                </el-button>
                <!-- <el-button type="primary" link icon="delete" size="small" @click="deleteRow(scope.row)">Delete
                </el-button> -->
              </template>
            </el-table-column>
          </el-table>
          <div class="gva-pagination">
            <el-pagination layout="total, sizes, prev, pager, next, jumper" :current-page="page" :page-size="pageSize"
              :page-sizes="[10, 30, 50, 100]" :total="total" @current-change="handleCurrentChange"
              @size-change="handleSizeChange" />
          </div>
        </el-tab-pane>
        <el-tab-pane name="Processing" label="Processing">
          <el-table ref="multipleTable" style="width: 100%" tooltip-effect="dark" :data="tableData" row-key="ID"
            @selection-change="handleSelectionChange">
            <!-- <el-table-column type="selection" width="55" /> -->
            <el-table-column align="left" label="ID" prop="ID" width="60" />
            <el-table-column align="left" label="Name" prop="customerObject.name" width="200" />
            <el-table-column align="left" label="Phone" prop="customerObject.phone" width="150" />
            <!-- <el-table-column align="left" label="Email" prop="customerObject.email" width="180" /> -->
            <el-table-column align="left" label="Qty" prop="quantity" width="50" />
            <el-table-column align="left" label="Bus" prop="busNumber" width="50" />
            <el-table-column align="left" label="Seats" prop="seats" width="120" />
            <el-table-column align="left" label="Total" prop="total" width="120">
              <template #default="scope">{{ formatCurrency(scope.row.total) }}</template>
            </el-table-column>
            <el-table-column align="left" label="Coupon Code" prop="couponCode" width="120" />
            <el-table-column align="left" label="Discount" prop="couponDiscount" width="120" />
            <el-table-column align="left" label="Tax" prop="tax" width="120">
              <template #default="scope">{{ formatCurrency(scope.row.tax) }}</template>
            </el-table-column>
            <el-table-column align="left" label="Payment" prop="invoicePayment" width="120">
              <template #default="scope">{{ formatCurrency(scope.row.invoicePayment) }}</template>
            </el-table-column>
            <!-- <el-table-column align="left" label="Source" prop="source" width="120" /> -->
            <!-- <el-table-column align="left" label="Notes" prop="notes" width="120" /> -->
            <!-- <el-table-column align="left" label="Status" prop="status" width="120" /> -->

            <el-table-column align="left" label="Booking Time" width="180">
              <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
            </el-table-column>
            <el-table-column align="left" label="Action">
              <template #default="scope">
                <el-button type="primary" link size="small" class="table-button"
                  @click="confirmCancelToCompleted(scope.row)">Confirm
                </el-button>
                <!-- <el-button type="primary" link icon="delete" size="small" @click="deleteRow(scope.row)">Delete
                </el-button> -->
              </template>
            </el-table-column>
          </el-table>
          <div class="gva-pagination">
            <el-pagination layout="total, sizes, prev, pager, next, jumper" :current-page="page" :page-size="pageSize"
              :page-sizes="[10, 30, 50, 100]" :total="total" @current-change="handleCurrentChange"
              @size-change="handleSizeChange" />
          </div>
        </el-tab-pane>
        <el-tab-pane name="Refund" label="Refund">
        </el-tab-pane>
        <el-tab-pane name="Reserve" label="Reserve">
        </el-tab-pane>
        <el-tab-pane name="Reserved" label="Reserved">
        </el-tab-pane>
      </el-tabs>
    </div>
    <div class="static-content-item">
      <el-divider direction="horizontal"></el-divider>
    </div>
    <div class="card-container">
      <el-card style="{width: 100% !important}" shadow="hover">
        <!-- <div slot="header" class="clear-fix">
          <span>Activity</span>
          <i class="float-right el-icon-arrow-down"></i>
        </div> -->
      </el-card>
    </div>
    <el-form-item label="Quick Note" prop="appointmentNote" class="label-center-align">
      <vue-editor v-model="appointmentData.appointmentNote"></vue-editor>
    </el-form-item>
    <div class="static-content-item">
      <el-button @click="updateAppointment" type="primary">Save Quick Note</el-button>
    </div>
  </el-form>

</template>

<script>
export default {
  name: 'EzyAppointmentDetail',
}
</script>

<script setup>
import { reactive, ref } from 'vue'
import { getDictFunc, formatDate, formatCurrency } from '@/utils/format'
import { VueEditor } from 'vue3-editor'
import { useRoute } from 'vue-router'
import { ElMessage, ElMessageBox } from 'element-plus'
import { useUserStore } from '@/pinia/modules/user'
import EzyMap from '@/components/ezyMap/index.vue'
import {
  updateEzyAppointment,
  findEzyAppointment,
} from '@/api/ezyAppointment'

import {
  createEzyOrders,
  deleteEzyOrders,
  deleteEzyOrdersByIds,
  updateEzyOrders,
  findEzyOrders,
  getEzyOrdersList,
  getEzyOrdersListByAppointment,
} from '@/api/ezyOrders'

const userStore = useUserStore()
const activeTab = ref('Completed')

const route = useRoute()
const vForm = ref()

const ezyMapRef = ref()
const appointmentData = ref({
  appointmentName: '',
  slug: '',
  singer: '',
  appointmentDate: null,
  startAt: null,
  endAt: null,
  appointmentNote: '',
  appointmentContent: null,
  stageId: '',
  branchId: '',
  stageMap: '',
  stageArea: '',
  disableIndex: [],
  totalSeat: 0,
  hideStageIndex: false,
  allowBus: false,
  status: 'publish',
  createdBy: 1,
  featuredImage: null,
  branch: 3,
},
)

const stageMapData = ref([])

const stageAreaData = ref([
  {
    area: '',
    seats: [],
    price: 0,
    description: '',
  },
])

const booked = ref([])

const rules = ref({})
const searchInfo = ref({ appointmentId: Number(route.params.id) })

const bookingForm = ref({
  total: 600000,
  invoicePayment: 600000,
  adminDiscount: 0,
  couponDiscount: 0,
  membershipDiscount: 0,
  tax: 0,
  customerName: "Hoàng bách",
  customerPhone: "0352974899",
  customerEmail: "bachth30@gmail.com"
})
const bookingFormRules = ref(
  {
    customerName: [{
      required: true,
      message: 'Customer name cannot be empty',
      trigger: ['input', 'blur'],
    }],
    customerPhone: [{
      required: true,
      message: 'Customer name cannot be empty',
      trigger: ['input', 'blur'],
    }],
    customerEmail: [{
      required: true,
      message: 'Email cannot be empty',
      trigger: ['input', 'blur'],
    }],
    // seats: [{
    //   required: true,
    //   message: 'Please select seats',
    //   trigger: ['input', 'blur'],
    // }],
    paymentType: [{
      required: true,
      message: 'Please select payment type',
      trigger: ['input', 'blur'],
    }],
  }
)
const vBookingForm = ref({})

const getEzyAppointmentById = async () => {
  let appointmentId = searchInfo.value.appointmentId
  var res = await findEzyAppointment({ ID: appointmentId })

  appointmentData.value = res.data.reezyAppointment
  if (appointmentData.value.stageMap) {
    let stageMapObject = JSON.parse(appointmentData.value.stageMap)
    stageMapData.value = stageMapObject
    // debugger;
    console.log(stageMapData.value)
  }

  if (appointmentData.value.stageArea) {
    var areaObject = JSON.parse(appointmentData.value.stageArea)
    stageAreaData.value = areaObject
    console.log(stageAreaData.value)
  }
  if (appointmentData.value.disableIndex) {
    appointmentData.value.disableIndex = JSON.parse(appointmentData.value.disableIndex)
  }
  console.log(appointmentData.value)

  ezyMapRef.value.data = stageAreaData.value
  ezyMapRef.value.initMap()
}

getEzyAppointmentById()

const updateAppointment = async () => {
  vForm.value?.validate(async (valid) => {
    if (!valid) return
    var slug = toSlug(appointmentData.value.appointmentName)
    appointmentData.value.slug = slug
    appointmentData.value.stageMap = JSON.stringify(getStageMap())
    appointmentData.value.stageArea = JSON.stringify(getStageArea())
    appointmentData.value.disableIndex = JSON.stringify(appointmentData.value.disableIndex)
    var res = await updateEzyAppointment(appointmentData.value)
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: 'Update successfullys',
      })
      // closeDialog()
      // getOrdersTableData()
    }
  })
}

const calculateInvoicePayment = () => {
  var adminDiscount = bookingForm.value.adminDiscount * 1;
  var tax = bookingForm.value.tax * 1;
  var total = bookingForm.value.total * 1;
  var membershipDiscount = bookingForm.value.membershipDiscount * 1;
  var couponDiscount = bookingForm.value.couponDiscount * 1;
  var invoicePayment = total - adminDiscount - couponDiscount - membershipDiscount + tax;
  bookingForm.value.invoicePayment = invoicePayment;
}

const createOrder = async () => {
  debugger;
  vBookingForm.value?.validate(async (valid) => {
    if (!valid) return
    bookingFormRules.value.seats = "A1,A2,A3"
    bookingForm.value.status = 'Admin'
    bookingForm.value.adminDiscount = bookingForm.value.adminDiscount * 1;
    bookingForm.value.total = 0;
    bookingForm.value.tax = 0;
    bookingForm.value.couponCode = null;
    bookingForm.value.couponDiscount = null;
    bookingForm.value.invoicePayment = 0;
    bookingForm.value.quantity = 3;
    bookingForm.value.total = 600000;
    bookingForm.value.createdBy = userStore.userInfo.ID;
    bookingForm.value.appointmentID = searchInfo.value.appointmentId;
    let res = await createEzyOrders(bookingForm.value)
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: 'Create order success'
      })
      // closeDialog()
      getOrdersTableData();
    }
  })
}

const handleTabClick = async (tab, event) => {
  activeTab.value = tab.props.name
  await getOrdersTableData()
}

// =========== TABLE DATA HANDLE ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
// const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
}

// 搜索
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  getOrdersTableData()
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getOrdersTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getOrdersTableData()
}

// 查询
const getOrdersTableData = async () => {
  var status = activeTab.value
  const table = await getEzyOrdersListByAppointment({
    appointmentID: searchInfo.value.appointmentId,
    status: status,
    page: page.value,
    pageSize: pageSize.value, ...searchInfo.value,
  })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
  console.log("get order table data");
  console.log(tableData.value)
}

getOrdersTableData()

const orderSelected = ref()

const showActionConfirm = (message, confirmCallback) => {
  ElMessageBox.confirm(
    message,
    'Warning',
    {
      confirmButtonText: 'OK',
      cancelButtonText: 'Cancel',
      type: 'warning',
    },
  )
    .then(() => {
      confirmCallback()
    })
    .catch(() => {

    })
}

const updateStatusForOrder = async (orderId, status) => {
  const res = await findEzyOrders({ ID: orderId })
  if (res.code === 0) {
    orderSelected.value = res.data.reezyOrders
    orderSelected.value.status = status
    orderSelected.value.lastActionBy = userStore.userInfo.ID
    orderSelected.value.lastActionTime = new Date()
  }
  var resUpdate = await updateEzyOrders(orderSelected.value)
  if (resUpdate.code === 0) {
    ElMessage({
      type: 'success',
      message: 'Update Successful',
    })
    getOrdersTableData()
  }
}

const confirmProcessingToCompleted = async (row) => {
  showActionConfirm('Confirm for this order from Processing to Completed', async () => {
    await updateStatusForOrder(row.ID, 'Completed')
  })
}

const confirmAdminBookingToCompleted = async (row) => {
  showActionConfirm('Confirm for this order from Admin Booking to Completed', async () => {
    await updateStatusForOrder(row.ID, 'Completed')
  })
}

const cancelFromCompleted = async (row) => {
  showActionConfirm('Cancel for this order from Completed', async () => {
    await updateStatusForOrder(row.ID, 'Cancel')
  })
}

const confirmCancelToCompleted = async (row) => {
  showActionConfirm('Confirm for this order from Cancel to Completed', async () => {
    await updateStatusForOrder(row.ID, 'Completed')
  })
}

const reserveFromCompleted = async (row) => {
  showActionConfirm('Confirm for this order from Cancel to Completed', async () => {
    await updateStatusForOrder(row.ID, 'Reserve')
  })
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


.padding-8 {
  padding: 8px;
}

.w100 {
  width: 100%;
}

.h800 {
  // height: 800px;
}

.el-select--large {
  width: 100%;
  line-height: 40px;
}

.form-item-row {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  padding: auto;
}
</style>