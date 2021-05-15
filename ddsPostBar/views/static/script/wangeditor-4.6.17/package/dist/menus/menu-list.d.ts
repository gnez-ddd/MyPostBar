/**
 * @description 所有菜单的构造函数
 * @author wangfupeng
 */
import Bold from './bold';
import Head from './head';
import Link from './link';
import Italic from './italic';
import Underline from './underline';
import StrikeThrough from './strike-through';
import FontStyle from './font-style';
import FontSize from './font-size';
import Justify from './justify';
import Quote from './quote';
import BackColor from './back-color';
import FontColor from './font-color';
import Video from './video';
import Image from './img';
import Indent from './indent';
import Emoticon from './emoticon';
import List from './list';
import LineHeight from './lineHeight';
import Undo from './undo';
import Redo from './redo';
import Table from './table';
import Code from './code';
import SplitLine from './split-line';
import Todo from './todo';
export declare type MenuListType = {
    [key: string]: any;
};
declare const _default: {
    bold: typeof Bold;
    head: typeof Head;
    italic: typeof Italic;
    link: typeof Link;
    underline: typeof Underline;
    strikeThrough: typeof StrikeThrough;
    fontName: typeof FontStyle;
    fontSize: typeof FontSize;
    justify: typeof Justify;
    quote: typeof Quote;
    backColor: typeof BackColor;
    foreColor: typeof FontColor;
    video: typeof Video;
    image: typeof Image;
    indent: typeof Indent;
    emoticon: typeof Emoticon;
    list: typeof List;
    lineHeight: typeof LineHeight;
    undo: typeof Undo;
    redo: typeof Redo;
    table: typeof Table;
    code: typeof Code;
    splitLine: typeof SplitLine;
    todo: typeof Todo;
};
export default _default;
