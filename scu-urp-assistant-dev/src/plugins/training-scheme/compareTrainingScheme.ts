// import {
//   showLoadingAnimation,
//   hideLoadingAnimation,
//   genFooterHTML
// } from './common'
// import {
//   TrainingSchemeBaseInfo,
//   TrainingSchemeYearInfo as SingleTrainingSchemeYearInfo,
//   TrainingSchemeCourseInfo as SingleTrainingSchemeCourseInfoBase,
//   TrainingScheme
// } from '@/store/types'
// import { getChineseNumber } from '@/helper/getter'
// import { Request, actions, state } from '@/store'
// import { emitDataAnalysisEvent } from '../data-analysis'
// import html2canvas from 'html2canvas'
// import { notifyError, Num } from '@/helper/util'

// let trainingSchemeList: TrainingScheme[]

// interface TrainingSchemeYearInfo {
//   name: string
//   children: TrainingSchemeSemesterInfo[]
// }

// interface TrainingSchemeSemesterInfo {
//   name: string
//   children: TrainingSchemeCourseInfo[]
// }

// interface TrainingSchemeCourseInfo extends SingleTrainingSchemeCourseInfoBase {
//   courseNumber1: string
//   courseNumber2: string
//   courseName1: string
//   courseName2: string
//   coursePropertyName1: string
//   coursePropertyName2: string
//   courseAttributes1: string[]
//   courseAttributes2: string[]
//   courseMajor1: string
//   courseMajor2: string
//   comparedResult: 0 | 1 | 2
//   [key: string]: string | string[] | number
// }

// function initDOM(root: HTMLElement): void {
//   const template = `
//       <div class="sua-container-compare-training-scheme">
//       </div>
//     `
//   $(root).append(template)
// }

// function genQueryHTML(): string {
//   const { gradeList, departmentList } = trainingSchemeList.reduce(
//     (acc, cur) => ({
//       gradeList: acc.gradeList.includes(cur.grade)
//         ? acc.gradeList
//         : acc.gradeList.concat(cur.grade),
//       departmentList: acc.departmentList.includes(cur.department)
//         ? acc.departmentList
//         : acc.departmentList.concat(cur.department)
//     }),
//     { gradeList: [] as string[], departmentList: [] as string[] }
//   )
//   return `
//       <div class="query-container">
//         <div id="query-major-1" class="row query-major">
//           <div class="col-xs-12 self-margin">
//             <h4 class="header smaller lighter grey">
//               <i class="ace-icon fa fa-search"></i>查询条件
//               <span class="right_top_oper">
//                 <button id="queryButton" title="查询" class="btn btn-info btn-xs btn-round">
//                   <i class="ace-con fa fa-search white bigger-120"></i> 查询
//                 </button>
//                 <button id="saveButton" title="导出长图" class="btn btn-success btn-xs btn-round">
//                   <i class="ace-con fa fa-cloud-download white bigger-120"></i> 导出长图
//                 </button>
//               </span>
//             </h4>
//             <div class="profile-user-info profile-user-info-striped self">
//               <div class="profile-info-row">
//                 <div class="profile-info-value">专业-1</div>
//                 <div class="profile-info-name">年级</div>
//                 <div class="profile-info-value">
//                   <select name="grade" id="grade" class="select form-control value_element">
//                     <option value="请选择年级">请选择年级</option>
//                     ${gradeList
//                       .sort(
//                         (a, b) =>
//                           Num(b.replace('级', '')) -
//                           Num(a.replace('级', ''))
//                       )
//                       .map(v => `<option value="${v}">${v}</option>`)
//                       .join('')}
//                   </select>
//                 </div>
//                 <div class="profile-info-name">院系</div>
//                 <div class="profile-info-value">
//                   <select name="department" id="department" class="select form-control value_element">
//                     <option value="请选择学院">请选择学院</option>
//                     ${departmentList
//                       .map(v => `<option value="${v}">${v}</option>`)
//                       .join('')}
//                   </select>
//                 </div>
//                 <div class="profile-info-name">专业</div>
//                 <div class="profile-info-value">
//                   <select name="major" id="major" class="form-control value_element">
//                     <option value="无">无</option>
//                   </select>
//                 </div>
//               </div>
//             </div>
//           </div>
//         </div>
//         <div id="query-major-2" class="row query-major">
//           <div class="col-xs-12 self-margin">
//             <div class="profile-user-info profile-user-info-striped self">
//               <div class="profile-info-row">
//                 <div class="profile-info-value">专业-2</div>
//                 <div class="profile-info-name">年级</div>
//                 <div class="profile-info-value">
//                   <select name="grade" id="grade" class="select form-control value_element">
//                     <option value="请选择年级">请选择年级</option>
//                     ${gradeList
//                       .sort(
//                         (a, b) =>
//                           Num(b.replace('级', '')) -
//                           Num(a.replace('级', ''))
//                       )
//                       .map(v => `<option value="${v}">${v}</option>`)
//                       .join('')}
//                   </select>
//                 </div>
//                 <div class="profile-info-name">院系</div>
//                 <div class="profile-info-value">
//                   <select name="department" id="department" class="select form-control value_element">
//                     <option value="请选择学院">请选择学院</option>
//                     ${departmentList
//                       .map(v => `<option value="${v}">${v}</option>`)
//                       .join('')}
//                   </select>
//                 </div>
//                 <div class="profile-info-name">专业</div>
//                 <div class="profile-info-value">
//                   <select name="major" id="major" class="form-control value_element">
//                     <option value="无">无</option>
//                   </select>
//                 </div>
//               </div>
//             </div>
//           </div>
//         </div>
//       </div>
//     `
// }

// function initQueryDOM(): void {
//   $('.sua-container-compare-training-scheme').append(genQueryHTML())
// }

// function updateMajorList(containerSelector: string): void {
//   const grade = $(`${containerSelector} #grade`).val()
//   const department = $(`${containerSelector} #department`).val()
//   const res = trainingSchemeList
//     .filter(v => v.grade === grade && v.department === department)
//     .map(v => `<option value="${v.majorId}">${v.majorName}</option>`)
//     .join('')
//   $(`${containerSelector} #major`)
//     .empty()
//     .append(res || `<option value="无">无</option>`)
// }

// function genInfoHTML(
//   info1: TrainingSchemeBaseInfo,
//   info2: TrainingSchemeBaseInfo
// ): string {
//   const preprocess = (info: TrainingSchemeBaseInfo): void => {
//     for (const key of Object.keys(info)) {
//       if (!info[key]) {
//         info[key] = '-'
//       }
//     }
//   }
//   ;[info1, info2].forEach(v => preprocess(v))
//   const genItemHTML = (
//     info1: TrainingSchemeBaseInfo,
//     info2: TrainingSchemeBaseInfo,
//     key: string
//   ): string =>
//     info1[key] === info2[key]
//       ? `<span class="item-value-same">${info1[key]}</span>`
//       : `
//           <div class="item-value-1">${info1.zym}（${info1.njmc}）：${info1[key]}</div>
//           <div class="item-value-2">${info2.zym}（${info2.njmc}）：${info2[key]}</div>
//         `
//   return `
//       <div class="info-container">
//         <div class="info-header row">
//           <div class="col-xs-12">
//             <h4 class="header smaller lighter grey">
//               <i class="fa fa-graduation-cap"></i> 方案计划信息比较
//             </h4>
//           </div>
//         </div>
//         <div class="info-header row">
//           <div class="col-xs-12">
//             <div class="info-tip">提示：在方案计划信息比较结果中，<span class="item-value-same">绿色</span>表示该项两个专业的值相同，<span class="item-value-1">红色</span>表示 <span class="item-value-1">专业-1</span> 的值，<span class="item-value-2">蓝色</span>表示 <span class="item-value-2">专业-2</span> 的值</div>
//           </div>
//         </div>
//         <div class="info-content row">
//           <div class="col-xs-12 col-md-4">
//             <table class="table table-bordered table-hover">
//               <tbody>
//                 <tr>
//                   <td>方案名称</td>
//                   <td>${genItemHTML(info1, info2, 'famc')}</td>
//                 </tr>
//                 <tr>
//                   <td>计划名称</td>
//                   <td>${genItemHTML(info1, info2, 'jhmc')}</td>
//                 </tr>
//                 <tr>
//                   <td>年级</td>
//                   <td>${genItemHTML(info1, info2, 'njmc')}</td>
//                 </tr>
//                 <tr>
//                   <td>院系名称</td>
//                   <td>${genItemHTML(info1, info2, 'xsm')}</td>
//                 </tr>
//                 <tr>
//                   <td>专业名称</td>
//                   <td>${genItemHTML(info1, info2, 'zym')}</td>
//                 </tr>
//                 <tr>
//                   <td>专业方向名称</td>
//                   <td>${genItemHTML(info1, info2, 'zyfxm')}</td>
//                 </tr>
//                 <tr>
//                   <td>学位</td>
//                   <td>${genItemHTML(info1, info2, 'xwm')}</td>
//                 </tr>
//               </tbody>
//             </table>
//           </div>
//           <div class="col-xs-12 col-md-4">
//             <table class="table table-bordered table-hover">
//               <tbody>
//                 <tr>
//                   <td>毕业类型</td>
//                   <td>${genItemHTML(info1, info2, 'bylxmc')}</td>
//                 </tr>
//                 <tr>
//                   <td>学制类型</td>
//                   <td>${genItemHTML(info1, info2, 'xzlxmc')}</td>
//                 </tr>
//                 <tr>
//                   <td>修读类型</td>
//                   <td>${genItemHTML(info1, info2, 'xdlxmc')}</td>
//                 </tr>
//                 <tr>
//                   <td>方案计划类型</td>
//                   <td>${genItemHTML(info1, info2, 'fajhlx')}</td>
//                 </tr>
//                 <tr>
//                   <td>开始学年</td>
//                   <td>${genItemHTML(info1, info2, 'xnmc')}</td>
//                 </tr>
//                 <tr>
//                   <td>学期类型</td>
//                   <td>${genItemHTML(info1, info2, 'xqlxm')}</td>
//                 </tr>
//                 <tr>
//                   <td>开始学期</td>
//                   <td>${genItemHTML(info1, info2, 'xqm')}</td>
//                 </tr>
//               </tbody>
//             </table>
//           </div>
//           <div class="col-xs-12 col-md-4">
//             <table class="table table-bordered table-hover">
//               <tbody>
//                 <tr>
//                   <td>要求总学分</td>
//                   <td>${genItemHTML(info1, info2, 'yqzxf')}</td>
//                 </tr>
//                 <tr>
//                   <td>课程总学分</td>
//                   <td>${genItemHTML(info1, info2, 'kczxf')}</td>
//                 </tr>
//                 <tr>
//                   <td>课程总门数</td>
//                   <td>${genItemHTML(info1, info2, 'kczms')}</td>
//                 </tr>
//                 <tr>
//                   <td>课程总学时</td>
//                   <td>${genItemHTML(info1, info2, 'kczxs')}</td>
//                 </tr>
//                 <tr>
//                   <td>学制类型</td>
//                   <td>${genItemHTML(info1, info2, 'xzlxmc')}</td>
//                 </tr>
//                 <tr>
//                   <td>培养目标</td>
//                   <td>${genItemHTML(info1, info2, 'pymb')}</td>
//                 </tr>
//                 <tr>
//                   <td>修读要求</td>
//                   <td>${genItemHTML(info1, info2, 'xdyq')}</td>
//                 </tr>
//               </tbody>
//             </table>
//           </div>
//           <div class="col-xs-12">
//             <table class="table table-bordered table-hover">
//               <tbody>
//                 <tr>
//                   <td>备注</td>
//                   <td>${genItemHTML(info1, info2, 'bz')}</td>
//                 </tr>
//               </tbody>
//             </table>
//           </div>
//         </div>
//       </div>
//     `
// }

// function genSchemeHTML(
//   list1: SingleTrainingSchemeYearInfo[],
//   list2: SingleTrainingSchemeYearInfo[]
// ): string {
//   const preprocess = (
//     list1: SingleTrainingSchemeYearInfo[],
//     list2: SingleTrainingSchemeYearInfo[]
//   ): Array<{
//     name: string
//     children: TrainingSchemeSemesterInfo[]
//   }> => {
//     const totalYear = Math.max(list1.length, list2.length)
//     const result = []
//     for (let i = 0; i < totalYear; i++) {
//       // 第一层循环，处理每个学年数据的合并
//       const yearItem1: TrainingSchemeYearInfo = (list1[
//         i
//       ] as TrainingSchemeYearInfo) || {
//         name: '',
//         children: [] as TrainingSchemeSemesterInfo[]
//       }
//       const yearItem2: TrainingSchemeYearInfo = (list2[
//         i
//       ] as TrainingSchemeYearInfo) || {
//         name: '',
//         children: [] as TrainingSchemeSemesterInfo[]
//       }
//       const yearItem = {
//         name:
//           yearItem1.name === yearItem2.name
//             ? yearItem1.name
//             : [yearItem1.name, yearItem2.name].filter(v => v).join(' / '),
//         children: [] as TrainingSchemeSemesterInfo[]
//       }
//       const totalSemester = Math.max(
//         yearItem1.children.length,
//         yearItem2.children.length
//       )
//       for (let j = 0; j < totalSemester; j++) {
//         // 第二层循环，处理每个学期数据的合并
//         const semesterItem1 = yearItem1.children[j] || {
//           name: '',
//           children: [] as TrainingSchemeCourseInfo[]
//         }
//         const semesterItem2 = yearItem2.children[j] || {
//           name: '',
//           children: [] as TrainingSchemeCourseInfo[]
//         }
//         const semesterItem = {
//           name:
//             semesterItem1.name === semesterItem2.name
//               ? semesterItem1.name
//               : [semesterItem1.name, semesterItem2.name]
//                   .filter(v => v)
//                   .join(' / '),
//           children: [] as TrainingSchemeCourseInfo[]
//         }
//         // 所有Key对应的值都一样
//         // const hasAllSameValues = (objs, keys) => {
//         //   for (const key of keys) {
//         //     if (!objs.reduce((acc, cur, i, src) => acc ? cur[key] === src[0][key] : acc, true)) {
//         //       return false
//         //     }
//         //   }
//         //   return true
//         // }
//         // 存在一个Key对应的值一样
//         const hasASameValue = (
//           objs: TrainingSchemeCourseInfo[],
//           keys: string[]
//         ): boolean => {
//           for (const key of keys) {
//             if (
//               objs.reduce(
//                 (acc, cur, i, src) => (acc ? cur[key] === src[0][key] : acc),
//                 true
//               )
//             ) {
//               return true
//             }
//           }
//           return false
//         }
//         const sameCourses = [] as unknown[]
//         const coursesOnly1 = [] as unknown[]
//         while (semesterItem1.children.length) {
//           const courseItem1 = semesterItem1.children.shift() as TrainingSchemeCourseInfo
//           let hasSameCourse = false
//           for (let m = 0; m < semesterItem2.children.length; m++) {
//             const courseItem2 = semesterItem2.children[m]
//             if (
//               hasASameValue(
//                 [courseItem1, courseItem2],
//                 ['courseName', 'courseNumber']
//               )
//             ) {
//               sameCourses.push({
//                 courseNumber1: courseItem1.courseNumber,
//                 courseNumber2: courseItem2.courseNumber,
//                 courseName1: courseItem1.courseName,
//                 courseName2: courseItem2.courseName,
//                 coursePropertyName1: courseItem1.coursePropertyName,
//                 coursePropertyName2: courseItem2.coursePropertyName,
//                 courseAttributes1: courseItem1.courseAttributes,
//                 courseAttributes2: courseItem2.courseAttributes,
//                 courseMajor1: courseItem1.courseMajor,
//                 courseMajor2: courseItem2.courseMajor,
//                 comparedResult: 0
//               } as TrainingSchemeCourseInfo)
//               semesterItem2.children.splice(m, 1)
//               hasSameCourse = true
//               break
//             }
//           }
//           if (!hasSameCourse) {
//             coursesOnly1.push(Object.assign(courseItem1, { comparedResult: 1 }))
//           }
//         }
//         const coursesOnly2 = semesterItem2.children.map(v =>
//           Object.assign(v, { comparedResult: 2 })
//         )
//         semesterItem.children = sameCourses.concat(
//           coursesOnly1,
//           coursesOnly2
//         ) as TrainingSchemeCourseInfo[]
//         yearItem.children.push(semesterItem)
//       }
//       result.push(yearItem)
//     }
//     return result
//   }

//   const courseItemTemplate = (
//     course: TrainingSchemeCourseInfo,
//     number: number
//   ): string => {
//     const getCourseItemClass = (course: TrainingSchemeCourseInfo): string => {
//       const classList = {
//         0: 'item-value-same',
//         1: 'item-value-1',
//         2: 'item-value-2'
//       } as {
//         [key: number]: string
//       }
//       return classList[course.comparedResult]
//     }
//     const courseItemInfoSecondary = (
//       course: TrainingSchemeCourseInfo
//     ): string => {
//       let courseNumberTemplate = ''
//       const getCoursePropertyNameClass = (coursePropertyName: string): string =>
//         `info-tag course-property-name${
//           coursePropertyName === '必修' ? ' required' : ''
//         }`
//       let coursePropertyNameTemplate = ''
//       let courseAttributesTemplate = ''
//       switch (course.comparedResult) {
//         case 0:
//           if (course.courseNumber1 === course.courseNumber2) {
//             courseNumberTemplate = `<div class="info-tag course-number">${course.courseNumber1}</div>`
//           } else {
//             courseNumberTemplate = `
//                 <div class="info-tag course-number"><span class="item-value-1">${course.courseMajor1}：</span>${course.courseNumber1}</div>
//                 <div class="info-tag course-number"><span class="item-value-2">${course.courseMajor2}：</span>${course.courseNumber2}</div>
//                `
//           }
//           if (course.coursePropertyName1 === course.coursePropertyName2) {
//             coursePropertyNameTemplate = course.coursePropertyName1
//               ? `<div class="${getCoursePropertyNameClass(
//                   course.coursePropertyName1
//                 )}">${course.coursePropertyName1}</div>`
//               : ''
//           } else {
//             coursePropertyNameTemplate = `
//                 ${
//                   course.coursePropertyName1
//                     ? `<div class="${getCoursePropertyNameClass(
//                         course.coursePropertyName1
//                       )}"><span class="item-value-1">${
//                         course.courseMajor1
//                       }：</span>${course.coursePropertyName1}</div>`
//                     : ''
//                 }
//                 ${
//                   course.coursePropertyName2
//                     ? `<div class="${getCoursePropertyNameClass(
//                         course.coursePropertyName2
//                       )}"><span class="item-value-2">${
//                         course.courseMajor2
//                       }：</span>${course.coursePropertyName2}</div>`
//                     : ''
//                 }
//               `
//           }
//           const sameAttributes = []
//           const attributesOnly1 = []
//           while (course.courseAttributes1.length) {
//             const attr = course.courseAttributes1.shift() as string
//             if (course.courseAttributes2.includes(attr)) {
//               course.courseAttributes2.splice(
//                 course.courseAttributes2.indexOf(attr),
//                 1
//               )
//               sameAttributes.push(attr)
//             } else {
//               attributesOnly1.push(
//                 `<span class="item-value-1">${course.courseMajor1}：</span>${attr}`
//               )
//             }
//           }
//           const attributesOnly2 = course.courseAttributes2.map(
//             v =>
//               `<span class="item-value-2">${course.courseMajor2}：</span>${v}`
//           )
//           courseAttributesTemplate = sameAttributes
//             .concat(attributesOnly1, attributesOnly2)
//             .map(v => `<div class="info-tag course-attribute">${v}</div>`)
//             .join('&nbsp;')
//           break
//         case 1:
//         case 2:
//           courseNumberTemplate = `<div class="info-tag course-number">${course.courseNumber}</div>`
//           coursePropertyNameTemplate = course.coursePropertyName
//             ? `<div class="${getCoursePropertyNameClass(
//                 course.coursePropertyName
//               )}">${course.coursePropertyName}</div>`
//             : ''
//           courseAttributesTemplate = course.courseAttributes
//             .map(v => `<div class="info-tag course-attribute">${v}</div>`)
//             .join('&nbsp;')
//           break
//       }
//       return `
//           <div class="info-secondary">
//             ${courseNumberTemplate}
//             ${coursePropertyNameTemplate}
//             ${courseAttributesTemplate}
//           </div>
//         `
//     }
//     const getCourseName = (course: TrainingSchemeCourseInfo): string => {
//       switch (course.comparedResult) {
//         case 0:
//           if (course.courseName1 === course.courseName2) {
//             return `<span class="item-value-same">${course.courseName1}</span>`
//           } else {
//             return `
//                 <span class="item-value-1">${course.courseName1}</span>
//                 <span>/</span>
//                 <span class="item-value-2">${course.courseName2}</span>
//               `
//           }
//         case 1:
//           return `<span class="item-value-1">${course.courseName}</span>`
//         case 2:
//           return `<span class="item-value-2">${course.courseName}</span>`
//       }
//     }
//     return `
//         <div class="course-item-wrapper">
//           <div class="course-item ${getCourseItemClass(course)}">
//             <div class="course-item-info">
//               <div class="info-primary">
//                 <div class="course-name">
//                   <div>${number}. ${getCourseName(course)}</div>
//                 </div>
//               </div>
//               ${courseItemInfoSecondary(course)}
//             </div>
//           </div>
//         </div>
//       `
//   }
//   const semesterItemTemplate = (
//     semester: TrainingSchemeSemesterInfo
//   ): string => `
//       <div class="semester-item">
//         <div class="semester-item-title">${semester.name}</div>
//         <div class="semester-item-content">
//           ${semester.children
//             .map((v, i) => courseItemTemplate(v, i + 1))
//             .join('')}
//         </div>
//       </div>
//     `

//   const yearItemTemplate = (
//     year: TrainingSchemeYearInfo,
//     grade: number
//   ): string => {
//     const yearNames = year.name.split(' / ')
//     let yearName
//     if (yearNames.length > 1) {
//       yearName = `<span class="item-value-1">${yearNames[0]}</span> / <span class="item-value-2">${yearNames[1]}</span>`
//     } else {
//       yearName = `<span class="item-value-same">${yearNames[0]}</span>`
//     }
//     return `
//         <div class="year-item">
//           <div class="year-item-title"><i class="fa fa-cubes" aria-hidden="true"></i> ${yearName}（${getChineseNumber(
//       grade
//     )}年级）</div>
//           <div class="year-item-content">
//             ${year.children
//               .map(v => semesterItemTemplate(v))
//               .join('<div class="semester-divider"></div>')}
//           </div>
//         </div>
//       `
//   }
//   return `
//       <div class="scheme-container">
//         <div class="row">
//           <div class="col-xs-12">
//             <h4 class="header smaller lighter grey">
//               <i class="fa fa-book"></i> 培养方案与指导性教学计划比较
//             </h4>
//           </div>
//         </div>
//         <div class="info-header row">
//           <div class="col-xs-12">
//             <div class="info-tip">提示：在培养方案与指导性教学计划比较结果中，<span class="item-value-same">绿色</span>表示该课程在两个专业中的课程名或课程号相同，<span class="item-value-1">红色</span>表示 <span class="item-value-1">专业-1</span> 独有的课程，<span class="item-value-2">蓝色</span>表示 <span class="item-value-2">专业-2</span> 独有的课程</div>
//           </div>
//         </div>
//         <div class="row">
//           <div class="col-xs-12">
//             <div class="scheme-wrapper">
//               ${preprocess(list1, list2)
//                 .map((v, i) => yearItemTemplate(v, i + 1))
//                 .join('')}
//             </div>
//           </div>
//         </div>
//       </div>
//     `
// }

// async function query(): Promise<void> {
//   const majorNumber1 = $('#query-major-1 #major').val()
//   const majorNumber2 = $('#query-major-2 #major').val()
//   if (majorNumber1 !== '无' && majorNumber2 !== '无') {
//     $('.program-plan-wrapper').remove()
//     showLoadingAnimation('.sua-container-compare-training-scheme')
//     try {
//       const [
//         { info: info1, list: list1 },
//         { info: info2, list: list2 }
//       ] = await Promise.all([
//         actions[Request.TRAINING_SCHEME](Num(majorNumber1)),
//         actions[Request.TRAINING_SCHEME](Num(majorNumber2))
//       ])
//       hideLoadingAnimation()
//       $('.sua-container-compare-training-scheme').append(
//         `<div class="program-plan-wrapper"></div>`
//       )
//       $('.program-plan-wrapper').append(genInfoHTML(info1, info2))
//       $('.program-plan-wrapper').append(genSchemeHTML(list1, list2))
//       $('.program-plan-wrapper').append(genFooterHTML())
//       $('.footer-container').hide()
//       const majorName1 = trainingSchemeList.filter(
//         ({ majorId }) => majorId === majorNumber1
//       )[0].majorName
//       const majorName2 = trainingSchemeList.filter(
//         ({ majorId }) => majorId === majorNumber2
//       )[0].majorName
//       const grade1 = $('#query-major-1 #grade').val()
//       const grade2 = $('#query-major-2 #grade').val()
//       emitDataAnalysisEvent('培养方案比较', '查询成功', {
//         专业代码: `${majorNumber1}-${majorNumber2}`,
//         专业名称: `${majorName1}（${grade1}）-${majorName2}（${grade2}）`
//       })
//     } catch (error) {
//       emitDataAnalysisEvent('培养方案比较', '数据获取失败')
//     }
//   }
// }

// function save(): void {
//   if ($('.program-plan-wrapper').length) {
//     $('.footer-container').show()
//     window.urp.alert('正在生成培养方案比较结果长图，请稍作等待')
//     setTimeout(async () => {
//       const canvas = await html2canvas($('.program-plan-wrapper')[0], {
//         useCORS: true
//       })
//       canvas.toBlob(blob => {
//         const majorName1 = trainingSchemeList.filter(
//           ({ majorId }) => majorId === Num($('#query-major-1 #major').val())
//         )[0].majorName
//         const majorName2 = trainingSchemeList.filter(
//           ({ majorId }) => majorId === Num($('#query-major-2 #major').val())
//         )[0].majorName
//         const grade1 = $('#query-major-1 #grade').val()
//         const grade2 = $('#query-major-2 #grade').val()
//         const department1 = $('#query-major-1 #department').val()
//         const department2 = $('#query-major-2 #department').val()
//         const e = document.createEvent('MouseEvents')
//         const a = document.createElement('a')
//         const filename = `【${department1}-${majorName1}（${grade1}）】与【${department2}-${majorName2}（${grade2}）】的培养方案比较结果`
//         a.download = filename
//         a.href = window.URL.createObjectURL(blob)
//         a.dataset.downloadurl = ['image/png', a.download, a.href].join(':')
//         e.initMouseEvent(
//           'click',
//           true,
//           false,
//           window,
//           0,
//           0,
//           0,
//           0,
//           0,
//           false,
//           false,
//           false,
//           false,
//           0,
//           null
//         )
//         a.dispatchEvent(e)
//         $('.footer-container').hide()
//         window.urp.alert('图片文件下载已启动，请注意保存哦')
//       })
//     }, 0)
//   } else {
//     window.urp.alert('请先按查询按钮，再保存结果哦')
//   }
// }

// function initEvents(): void {
//   $('#query-major-1 #grade').change(() => {
//     updateMajorList('#query-major-1')
//   })
//   $('#query-major-1 #department').change(() => {
//     updateMajorList('#query-major-1')
//   })
//   $('#query-major-2 #grade').change(() => {
//     updateMajorList('#query-major-2')
//   })
//   $('#query-major-2 #department').change(() => {
//     updateMajorList('#query-major-2')
//   })
//   $('#queryButton').click(query)
//   $('#saveButton').click(save)
// }

// async function selectSelfMajorAndQuery(): Promise<void> {
//   const selfMajorNumber = state.user.programPlanNumber
//   const selfSchemeInfo = trainingSchemeList.filter(
//     v => Num(v.majorId) === selfMajorNumber
//   )[0]
//   $('#query-major-1 #grade').val(selfSchemeInfo.grade)
//   $('#query-major-2 #grade').val(selfSchemeInfo.grade)
//   $('#query-major-1 #department').val(selfSchemeInfo.department)
//   $('#query-major-2 #department').val(selfSchemeInfo.department)
//   updateMajorList('#query-major-1')
//   updateMajorList('#query-major-2')
//   $('#query-major-1 #major').val(selfSchemeInfo.majorId)
//   $('#query-major-2 #major').val(selfSchemeInfo.majorId)
// }

// export async function render(root: HTMLElement): Promise<void> {
//   initDOM(root)
//   showLoadingAnimation('.sua-container-compare-training-scheme')
//   try {
//     trainingSchemeList = await actions[Request.TRAINING_SCHEME_LIST]()
//     hideLoadingAnimation()
//     initQueryDOM()
//     initEvents()
//     selectSelfMajorAndQuery()
//   } catch (error) {
//     const title = '培养方案列表数据获取失败'
//     notifyError(error, title)
//     emitDataAnalysisEvent('培养方案比较', '培养方案列表数据获取失败')
//   }
// }
