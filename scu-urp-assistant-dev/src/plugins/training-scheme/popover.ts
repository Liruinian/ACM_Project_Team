import { actions, Request, state } from '@/store'
import { CourseScheduleInfo } from '@/store/types'

async function showCourseSchedulePop(
  semester: string,
  courseName: string,
  courseNumber: string
): Promise<void> {
  const { response, sequence } = await actions[Request.COURSE_SCHEDULE](
    semester,
    courseName,
    courseNumber
  )
  let titleText = ''
  let templateHTML = ''
  const $pop = $('.course-info-popover')
  const $popTitle = $pop.children('.ci-popover-title')
  const $popContent = $pop.children('.ci-popover-content')
  if (sequence) {
    if (sequence.length) {
      titleText = `${courseName}（${courseNumber}）- 共${sequence.length}个课序号`
      const genRowsHTML = (sequence: CourseScheduleInfo[]): string =>
        sequence
          .map(
            ({
              courseSequenceNumber,
              courseDeptName,
              credit,
              courseCategoryName,
              examTypeName,
              courseTeacher,
              courseTime,
              campusName,
              classCapacityRemaining,
              courseRegNote
            }) =>
              `<tr>
                <td>${courseSequenceNumber}</td>
                <td>${courseDeptName}</td>
                <td>${credit}</td>
                <td>${courseCategoryName}</td>
                <td>${examTypeName}</td>
                <td>${courseTeacher}</td>
                <td>${courseTime}</td>
                <td>${campusName}</td>
                <td>${classCapacityRemaining}</td>
                <td>${courseRegNote}</td>
              </tr>`
          )
          .join('')
      templateHTML = `
          <table class="table table-striped table-bordered table-hover">
            <thead>
              <tr>
                <th>课序号</th>
                <th>开课院系</th>
                <th>学分</th>
                <th>课程类别</th>
                <th>考试类型</th>
                <th>教师</th>
                <th>时间</th>
                <th>地点</th>
                <th>课余量</th>
                <th>备注</th>
              </tr>
            </thead>
            <tbody>
              ${genRowsHTML(sequence)}
            </tbody>
          </table>
        `
    } else {
      titleText = `${courseName}（${courseNumber}）`
      templateHTML = `<p style="color: #CB1B45;">【${courseName}（${courseNumber}）】并未在本学期开课</p>`
    }
  } else {
    if (response === 'network_error') {
      titleText = `${courseName}（${courseNumber}）`
      templateHTML = `<p style="color: #CB1B45;">因网络原因，【${courseName}（${courseNumber}）】的课程信息查询失败，请稍后再试</p>`
    }
  }
  if (titleText) {
    $popTitle.text(titleText)
  }
  if (templateHTML) {
    $popContent.html(templateHTML)
  }
}

export function initCourseInfoPopover(): void {
  const { currentSemesterNumber } = state.basic
  // 教务系统课程信息频繁查询的阈值
  const initDOM = function(
    element: HTMLElement,
    courseName: string,
    courseNumber: string
  ): void {
    if ($('.course-info-popover').length) {
      $('.course-info-popover').remove()
    }
    const loadingTips = `( º﹃º ) 兆基祈祷中……`
    $(element).append(`
      <div class="course-info-popover">
        <div class="ci-popover-title">${courseName}（${courseNumber}）</div>
        <div class="ci-popover-content">${loadingTips}</div>
        <div class="ci-popover-arrow"></div>
      </div>
    `)
    const $pop = $('.course-info-popover')
    const { left } = $pop.offset() || { left: 0 }
    const popWidth = $pop.width() || 0
    const bodyWidth = $('body').width() || 0
    if (left + popWidth > bodyWidth) {
      $pop.offset({ left: bodyWidth - popWidth - 80 })
    } else if (left < 0) {
      $pop.offset({ left: 50 })
    }
    $(element).mouseleave(function() {
      $pop.remove()
    })
  }
  $('.course-item-info').click(async function() {
    const $courseItem = $(this).parent()
    const courseName = $courseItem.data('course-name')
    const courseNumber = $courseItem.data('course-number')
    initDOM($courseItem[0], courseName, courseNumber)
    showCourseSchedulePop(currentSemesterNumber, courseName, courseNumber)
  })
}
