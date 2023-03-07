import Vue from 'vue'
import {
  Button,
  Switch,
  Tag,
  Alert,
  Loading,
  Message,
  MessageBox,
  Notification,
  Rate,
  Divider,
  Tooltip,
  CheckboxGroup,
  Checkbox,
  CheckboxButton
} from 'element-ui'
import JsonViewer from 'vue-json-viewer'
import { routeTrigger } from '@/helper/util'
import { Breadcrumbs, SUAPluginMenu } from './types'
import { addRoute, getCurrentRouteParams, RouteConfig, router } from './router'

/**
 * 加载 Vue 组件
 *
 */
export const loadElementUI = (): void => {
  // 导入 Element-UI 的样式
  $('head').append(
    '<link rel="stylesheet" href="//cdn.staticfile.org/element-ui/2.15.6/theme-chalk/index.min.css"></link>'
  )
  // 导入 Element-UI 组件
  Vue.use(Button)
  Vue.use(Switch)
  Vue.use(Tag)
  Vue.use(Alert)
  Vue.use(Loading.directive)
  Vue.use(Rate)
  Vue.use(Divider)
  Vue.use(Tooltip)
  Vue.use(CheckboxGroup)
  Vue.use(Checkbox)
  Vue.use(CheckboxButton)
  Vue.prototype.$loading = Loading.service
  Vue.prototype.$msgbox = MessageBox
  Vue.prototype.$alert = MessageBox.alert
  Vue.prototype.$confirm = MessageBox.confirm
  Vue.prototype.$prompt = MessageBox.prompt
  Vue.prototype.$notify = Notification
  Vue.prototype.$message = Message
  // 导入其他组件
  Vue.use(JsonViewer)
}

/**
 * 加载全局样式表
 *
 */
export const loadGlobalStyle = (): void => {
  const globalStyle = require('@/core/global.scss').toString()
  $('head').append(`
    <style type="text/css">
      ${globalStyle}
    </style>
  `)
}

export const loadRouteConfig = (routeConfig: RouteConfig): void => {
  addRoute(routeConfig)
  if (routeTrigger(routeConfig.path)) {
    router.push(routeConfig.path, {
      params: getCurrentRouteParams()
    })
  }
}

const changeMenuAndBreadcrumbs = (
  $rootMenu: JQuery<HTMLElement>,
  $menu: JQuery<HTMLElement>,
  $menuItem: JQuery<HTMLElement>,
  rootMenuId: string,
  breadcrumbs: Breadcrumbs,
  menuId: string,
  id: string
) => {
  const changeMenu = (
    $rootMenu: JQuery<HTMLElement>,
    $menu: JQuery<HTMLElement>,
    $menuItem: JQuery<HTMLElement>
  ) => {
    $('.hsub').removeClass('open')
    $('.submenu').css('display', 'none')
    $('.submenu>li').removeClass('active')
    $('.submenu>li>a>.menu-icon').remove()
    $rootMenu.parent().addClass('open')
    $rootMenu.css('display', 'block')
    $menu.parent().addClass('open')
    $menu.css('display', 'block')
    $menuItem.addClass('active')
    $menuItem.find('a').prepend("<i class='menu-icon fa fa-caret-right'></i>")
  }

  const showBreadcrumbs = (
    rootMenuId: string,
    breadcrumbs: Breadcrumbs,
    menuId: string,
    id: string
  ) => {
    const $breadcrumbs = $('.main-content>.breadcrumbs>ul.breadcrumb')
    $breadcrumbs.empty().append(`
    <li onclick="javascript:window.location.href='/'" style="cursor:pointer;">
      <i class="ace-icon fa fa-home home-icon"></i>
      首页
    </li>
    <li class="active" onclick="ckickTopMenu(this);return false;" id="firmenu" menuid="${rootMenuId}">${breadcrumbs[0]}</li>
    <li class="active" onclick="ckickTopMenu(this);return false;" id="secmenu" menuid="${menuId}">${breadcrumbs[1]}</li>
    <li class="active" onclick="ckickTopMenu(this);return false;" id="lastmenu" menuid="${id}">${breadcrumbs[2]}</li>
  `)
  }

  changeMenu($rootMenu, $menu, $menuItem)
  showBreadcrumbs(rootMenuId, breadcrumbs, menuId, id)
}

/**
 * 初始化插件菜单
 *
 * @param {SUAPluginMenu} menu 插件的菜单配置对象
 */
export const loadMenu = (menu: SUAPluginMenu): void => {
  const { rootMenuId, rootMenuName, id: menuId, name: menuName } = menu
  let { item: items } = menu
  if (!Array.isArray(items)) {
    items = [items]
  }
  // 将侧边栏置顶，避免在手机上使用时，顶栏因高度增加挡住侧边栏。
  const $sidebar = $('#sidebar')
  $sidebar.css('z-index', 9999)

  const $rootMenuList = $('#menus')
  // 检查根菜单是否存在，如不存在则新建
  if (!$rootMenuList.children(`li#${rootMenuId}`).length) {
    $rootMenuList.prepend(`
      <li class="hsub sua-menu-list" id="${rootMenuId}" onclick="rootMenuClick(this);">
        <a href="#" class="dropdown-toggle">
          <i class="menu-icon fa fa-gavel"></i>
          <span class="menu-text">${rootMenuName}</span>
          <b class="arrow fa fa-angle-down"></b>
        </a>
        <b class="arrow"></b>
        <ul class="submenu nav-hide" onclick="stopHere();" style="display: none;">
        </ul>
      </li>
    `)
  }
  const $rootMenu = $rootMenuList.find(`li#${rootMenuId}>ul.submenu`)
  // 检查菜单是否存在，如不存在则新建
  if (!$rootMenu.children(`li#${menuId}`).length) {
    $rootMenu.append(`
      <li class="hsub sua-menu" id="${menuId}">
        <a href="#" class="dropdown-toggle">
          <i class="menu-icon fa fa-caret-right"></i>${menuName}
          <b class="arrow fa fa-angle-down"></b></a>
        <b class="arrow"></b>
        <ul class="submenu nav-show" style="display: none;">
        </ul>
      </li>
    `)
  }
  const $menu = $rootMenu.find(`li#${menuId}>ul.submenu`)
  items.forEach(({ name, route }) => {
    const id = `menu-item-${name}`
    $menu.append(`
      <li class="sua-menu-item" id="${id}">
        <a href="#">&nbsp;&nbsp; ${name}</a>
        <b class="arrow"></b>
      </li>
    `)
    const $menuItem = $menu.children(`#${id}`)
    const breadcrumbs: Breadcrumbs = [rootMenuName, menuName, name]
    $menuItem.click(() => {
      changeMenuAndBreadcrumbs(
        $rootMenu,
        $menu,
        $menuItem,
        rootMenuId,
        breadcrumbs,
        menuId,
        id
      )
      router.push(route)
    })
    if (routeTrigger(route)) {
      changeMenuAndBreadcrumbs(
        $rootMenu,
        $menu,
        $menuItem,
        rootMenuId,
        breadcrumbs,
        menuId,
        id
      )
    }
  })
}
