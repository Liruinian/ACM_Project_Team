<template lang="pug">
.sua-container-gpa-calculator
  Loading(v-if='!loadingIsDone')
  el-alert(
    v-if='loadingIsDone',
    v-for='(v, i) in alerts',
    :key='i',
    :title='v.title',
    :type='v.type',
    :closable='v.closable',
    :close-text='v.closeText',
    style='margin-bottom: 10px'
  )
  TotalTranscript(
    v-if='loadingIsDone && hasNoError && records.length > 1',
    :semestersQuantity='records.length',
    :courses='allCourses',
    :selectedCourses='allSelectedCourses',
    @selectAllCourses='selectAllCourses()',
    @unselectAllCourses='unselectAllCourses()',
    @selectCompulsoryCourses='selectCompulsoryCourses()',
    @selectMinorCourses='selectMinorCourses()'
  )
  .gpa-st-container.row(v-if='loadingIsDone && hasNoError')
    SemesterTranscript(
      v-for='(semesterItem, semesterIndex) in records',
      :key='semesterIndex',
      :type='type',
      :semester='semesterItem.semester',
      :courses='semesterItem.courses'
    )
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'
import { SemesterScoreRecord, CourseScoreRecord } from '@/plugins/score/types'
import Loading from '@/plugins/common/components/Loading.vue'
import TotalTranscript from './components/TotalTranscript.vue'
import SemesterTranscript from './components/SemesterTranscript/SemesterTranscript.vue'
import { emitDataAnalysisEvent } from '../data-analysis'
import { getSelectedCourses } from '@/plugins/score/utils'
import { convertCourseScoreInfoListToScoreRecords } from '@/helper/converter'
// import * as ueip from '@/plugins/user-experience-improvement-program'
import {
  requestAllPassingScores,
  requestThisTermCourseScoreInfoList
} from '@/store/actions/request'
import { notifyError } from '@/helper/util'

@Component({
  components: { Loading, SemesterTranscript, TotalTranscript }
})
export default class GPACalculator extends Vue {
  @Prop({
    type: String,
    required: true
  })
  type!: 'full' | 'basic' | 'compact'

  loadingIsDone = false
  records: SemesterScoreRecord[] = []
  alerts: {
    title: string
    type?: 'success' | 'info' | 'warning' | 'error'
    closable?: boolean
    closeText?: string
  }[] = []

  get hasNoError(): boolean {
    return this.alerts.every(v => v.type !== 'error')
  }

  get allCourses(): CourseScoreRecord[] {
    return this.records.reduce(
      (acc, cur) => acc.concat(cur.courses),
      [] as CourseScoreRecord[]
    )
  }

  get allSelectedCourses(): CourseScoreRecord[] {
    return getSelectedCourses(this.allCourses)
  }

  selectAllCourses(): void {
    this.allCourses.forEach(v => (v.selected = true))
  }

  unselectAllCourses(): void {
    this.allCourses.forEach(v => (v.selected = false))
  }

  selectCompulsoryCourses(): void {
    this.allCourses.forEach(v => (v.selected = v.coursePropertyName === '??????'))
  }

  selectMinorCourses(): void {
    this.allCourses.forEach(v => (v.selected = v.coursePropertyName === '??????'))
  }

  async created(): Promise<void> {
    try {
      const records =
        this.type === 'full'
          ? convertCourseScoreInfoListToScoreRecords(
              await requestThisTermCourseScoreInfoList()
            )
          : convertCourseScoreInfoListToScoreRecords(
              await requestAllPassingScores()
            )
      this.records = records
      this.loadingIsDone = true
      // ueip.sendStudentCourseScorePublicList(records)
      switch (this.type) {
        case 'compact':
          emitDataAnalysisEvent('???????????????????????????', '????????????')
          break
        case 'basic':
          emitDataAnalysisEvent('?????????????????????', '????????????')
          break
        case 'full':
          emitDataAnalysisEvent('??????????????????', '????????????')
          break
      }
    } catch (error) {
      let title = '[??????????????????] '
      const message: string = (error as Error).message
      switch (this.type) {
        case 'compact':
          title += '???????????????????????????'
          emitDataAnalysisEvent('???????????????????????????', '????????????')
          break
        case 'basic':
          title += '?????????????????????'
          emitDataAnalysisEvent('?????????????????????', '????????????')
          break
        case 'full':
          title += '??????????????????'
          emitDataAnalysisEvent('??????????????????', '????????????')
          break
      }
      notifyError(message, title)
      this.alerts.push({
        title: message,
        type: 'error',
        closable: false
      })
      this.loadingIsDone = true
    }
    if (this.type === 'full') {
      this.alerts.push({
        title:
          '?????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????API??????????????????????????????????????????????????????????????????????????????????????????????????????',
        type: 'warning',
        closable: true
      })
    }
  }
}
</script>

<style lang="scss">
.gpa-st-container {
  display: flex;
  flex-wrap: wrap;
}

.gpa-st-select-all-btn,
.gpa-tt-select-all-btn {
  position: relative;
  top: 2.5px;
  float: right;
}

.gpa-st-cancel-btn,
.gpa-tt-cancel-btn {
  position: relative;
  top: 2.5px;
  float: right;
}

.gpa-st-tag,
.gpa-tt-tag {
  cursor: pointer;
}

.gpa-info-badge {
  cursor: pointer;
  position: relative;
  top: -7.5px;
}
</style>
