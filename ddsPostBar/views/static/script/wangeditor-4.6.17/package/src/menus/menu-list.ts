/**
 * @description 所有菜单的构造函数
 * @author wangfupeng
 */

import Bold from './bold'
import Head from './head'
import Link from './link'
import Italic from './italic'
import Underline from './underline'
import StrikeThrough from './strike-through'
import FontStyle from './font-style'
import FontSize from './font-size'
import Justify from './justify'
import Quote from './quote'
import BackColor from './back-color'
import FontColor from './font-color'
import Video from './video'
import Image from './img'
import Indent from './indent'
import Emoticon from './emoticon'
import List from './list'
import LineHeight from './lineHeight'
import Undo from './undo'
import Redo from './redo'
import Table from './table'
import Code from './code'
import SplitLine from './split-line'
import Todo from './todo'

export type MenuListType = {
    [key: string]: any
}

export default {
    bold: Bold,
    head: Head,
    italic: Italic,
    link: Link,
    underline: Underline,
    strikeThrough: StrikeThrough,
    fontName: FontStyle,
    fontSize: FontSize,
    justify: Justify,
    quote: Quote,
    backColor: BackColor,
    foreColor: FontColor,
    video: Video,
    image: Image,
    indent: Indent,
    emoticon: Emoticon,
    list: List,
    lineHeight: LineHeight,
    undo: Undo,
    redo: Redo,
    table: Table,
    code: Code,
    splitLine: SplitLine,
    todo: Todo,
}
