// 数据统计插件
import pack from '@/../package.json'
import { SUAPlugin } from '@/core/types'
import { isDev, isVersion } from '@/helper/judger'
import { Logger } from '@/helper/logger'
import { getPluginIcon } from '@/helper/getter'
import { getVersionName } from '@/helper/info'

const { version } = pack

interface TalkingDataEventParams {
  EventId: string
  Label?: string
  MapKv?: Record<string, string>
}

const queue: TalkingDataEventParams[] = []

function emitDataAnalysisEvent(
  EventId: string,
  Label?: string,
  MapKv?: Record<string, string>
): void {
  queue.push({ EventId, Label, MapKv })
}

export { emitDataAnalysisEvent }

export const DataAnalysis: SUAPlugin = {
  name: 'data-analysis',
  displayName: '数据统计',
  icon: getPluginIcon('data-analysis'),
  isNecessary: false,
  defaultEnabledState: true,
  brief:
    '为开发者提供诸如「助手版本」、「程序启动次数」、「功能使用频率」、「程序是否出错」等匿名统计数据，帮助开发者了解程序的运行情况。',
  pathname: true,
  task() {
    if (window.TDAPP && queue.length) {
      const e = queue.shift()
      if (e) {
        window.TDAPP.onEvent(e.EventId, e.Label, e.MapKv)
      }
    }
  },
  init() {
    if (process.env.NODE_ENV !== 'development') {
      const APP_ID = '36482C98B3E94A4D93A0C66E43702C77'
      const script = document.createElement('script')
      const versionName = getVersionName()
      const src = `https://jic.talkingdata.com/app/h5/v1?appid=${APP_ID}&vn=${versionName}&vc=${version}`
      script.setAttribute('src', src)
      document.getElementsByTagName('head')[0].appendChild(script)
      script.onload = (): void => {
        Logger.info(`Data Analysis Plugin Loaded Successfully: ${versionName}`)
      }
    }
  }
}
