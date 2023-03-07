// 成绩相关功能插件
import Vue, { VNode } from 'vue'
import GPACalculatorApp from './GPACalculator.vue'
import GPACalculatorWidgetApp from './GPACalculatorWidget.vue'
import ExpectedGradesEstimationApp from './ExpectedGradesEstimation.vue'
import { emitDataAnalysisEvent } from '../data-analysis'
import { getPluginIcon } from '@/helper/getter'
import { SUAPlugin } from '@/core/types'
import { RouteHookNextFunction } from '@/core/router'

const menu = [
  {
    rootMenuId: 'sua-menu-list',
    rootMenuName: 'SCU URP 助手',
    id: 'menu-utility-tools',
    name: '实用工具',
    item: {
      name: '均分绩点计算器',
      route: 'utility_tools/gpa_calculator'
    }
  },
  {
    rootMenuId: 'sua-menu-list',
    rootMenuName: 'SCU URP 助手',
    id: 'menu-utility-tools',
    name: '实用工具',
    item: {
      name: '预期成绩估计',
      route: 'utility_tools/expected_grades_estimation'
    }
  },
  {
    rootMenuId: 'sua-menu-list',
    rootMenuName: 'SCU URP 助手',
    id: 'menu-advanced-query',
    name: '高级查询',
    item: {
      name: '成绩信息查询',
      route: 'advanced_query/scores_information'
    }
  }
]

export const Score: SUAPlugin = {
  name: 'score',
  displayName: '成绩相关工具',
  icon: getPluginIcon('score'),
  isNecessary: false,
  defaultEnabledState: true,
  brief:
    '让您直接看到全部均分、全部绩点与必修均分、必修绩点，还可自由地选择课程进行计算，并方便地估计预期成绩。还可以在每年的出分季帮助您查询到本学期课程成绩的最高分、最低分、平均分和名次。',
  pathname: ['/', '/index.jsp'],
  route: [
    {
      path: 'utility_tools/gpa_calculator',
      component: GPACalculatorApp,
      componentOptions: {
        props: {
          type: 'basic'
        }
      }
    },
    {
      path: 'utility_tools/expected_grades_estimation',
      component: ExpectedGradesEstimationApp
    },
    {
      path: 'advanced_query/scores_information',
      component: GPACalculatorApp,
      componentOptions: {
        props: {
          type: 'full'
        }
      },
      beforeEnter(next: RouteHookNextFunction): void {
        window.urp.confirm(
          `<p style="font-weight: 700; color: red;">警告：</p>
          <p style="text-indent: 2em;">该页面展示的部分敏感数据（最高分、平均分、最低分、名次）调用了综合教务系统<span style="color: red;">【未公开】的接口</span>，如果综合教务系统关闭了该接口，那么这个功能就报废了，我们将无法再获取到这些教务系统屏蔽的数据！</p>
          <p style="text-indent: 2em;">因此，如果您要用本页面展示的这些敏感数据和您的任课老师沟通，我希望您可以<span style="color: red;">不要透露该插件的存在</span>，只是说这些敏感数据是您私下联系同学们询问成绩，从而获得的调查结果！</p>
          <p style="text-indent: 2em;">否则，老师一旦和教务处反映，这个数据获取接口就有<span style="color: red;">被关闭</span>的风险！</p>`,
          async (res: boolean) => {
            if (res) {
              next()
            } else {
              next(
                new Error(
                  '很抱歉，因为您拒绝了使用协议，SCU URP 助手 无法显示您希望看到的数据。'
                )
              )
              emitDataAnalysisEvent('成绩信息查询', '拒绝使用协议')
            }
          }
        )
      }
    }
  ],
  init() {
    $('.page-content').append(
      `<div class="row"><div class="sua-widget-container-gpa-calculator"></div></div>`
    )
    new Vue({
      render: (h): VNode => h(GPACalculatorWidgetApp)
    }).$mount('.sua-widget-container-gpa-calculator')
  },
  menu
}
