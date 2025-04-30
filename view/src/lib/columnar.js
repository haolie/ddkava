"use strict";
var columnar = function columnar(opt) {
  this.opt = opt;
  this.node = null;
  this.canvas = null;
  this.ctx = null;
  this.width = 0;
  this.height = 16;
  this.minHeight = 2;
  this.completedHeight = 0;
  this.addStyle();
  this.render();
};

columnar.prototype = {
  addStyle: function addStyle() {
    this.checkData();
    var node = document.querySelector('#columnar_custom_components');
    if (node) return;
    var style = document.createElement('style');
    style.id = 'columnar_custom_components';
    style.innerText = "\n      .columnar_custom_dom {\n        width: 100%;\n        height: 100%;\n        position: relative;\n        z-index: 1;\n        color: #fff;\n        text-align: center;\n        font-size: 12px;\n        box-sizing: border-box;\n      }\n      .columnar_custom_dom ul {\n        margin: 0;\n        padding: 0;\n      }\n      .columnar_custom_dom li {\n        list-style: none;\n      }\n      \n      .columnar_custom_dom em {\n        font-size: 20px;\n      }\n      \n      .total_target {\n        position: absolute;\n        z-index: 4;\n        width: 8rem;\n        height: 2rem;\n        left: 50%;\n        top: 1rem;\n        margin-left: -4rem;\n        line-height: 1.2rem;\n      }\n      \n      .columnar_total {\n        position: absolute;\n        z-index: 1;\n        width: 8rem;\n        height: 16rem;\n        left: 50%;\n        top: 4rem;\n        margin-left: -4rem;\n        background-image: linear-gradient(90deg, rgba(255, 255, 255, 0.1) 0%, rgba(255, 255, 255, 0.15) 25%, rgba(255, 255, 255, 0.15) 35%, rgba(255, 255, 255, 0.05) 60%, rgba(255, 255, 255, 0.05) 85%, rgba(255, 255, 255, 0.1) 100%);\n        line-height: 50px;\n      }\n      \n      .columnar_total_after {\n        position: absolute;\n        z-index: 2;\n        width: 8rem;\n        height: 2rem;\n        left: 50%;\n        top: 2rem;\n        margin-left: 4.5rem;\n        color: #FF4873;\n      }\n      \n      .columnar_total_after_border {\n        position: absolute;\n        z-index: 2;\n        display: block;\n        width: 6px;\n        height: 100%;\n        border: 1px solid #FF4873;\n        border-left: none;\n        top: 0;\n        left: 0;\n      }\n      \n      .columnar_total_min_height {\n        margin-top: -1rem;\n      }\n      \n      .columnar_total_min_height>.columnar_total_after_border {\n        margin-top: 1rem;\n      }\n      \n      .columnar_total_span {\n        position: absolute;\n        z-index: 4;\n        width: 8rem;\n        height: 16rem;\n        left: 50%;\n        top: 5rem;\n        margin-left: -4rem;\n        min-height: 1rem;\n      }\n      \n      .columnar_total_ellipse {\n        position: absolute;\n        z-index: 2;\n        width: 8rem;\n        height: 20px;\n        left: 50%;\n        top: 4rem;\n        margin-left: -4rem;\n        margin-top: -10px;\n        border-radius: 50% / 50%;\n        background-image: linear-gradient(90deg, #1C2E4B 0%, #3F5771 100%);\n        box-shadow: inset 1px -1px 1px rgba(255, 255, 255, 0.04);\n      }\n      \n      .columnar_ywc {\n        position: absolute;\n        z-index: 2;\n        width: 8rem;\n        height: 16rem;\n        border: 1px solid #ccc;\n        left: 50%;\n        top: 20px;\n        margin-left: -4rem;\n        background-image: linear-gradient(90deg, #254162 0%, #3E5875 10%, #4A6179 70%, #4A6179 70%);\n      }\n      \n      .columnar_fwc_half {\n        position: absolute;\n        z-index: 1;\n        width: 8rem;\n        height: 20px;\n        left: 50%;\n        top: 20rem;\n        margin-left: -4rem;\n        margin-top: -10px;\n        border-radius: 50% / 50%;\n        background-image: linear-gradient(90deg, #C8A325 0%, #EFD05F 25%, #EFD05F 35%, #DF8343 60%, #DF8343 85%, #DD9B82 100%);\n      }\n      \n      .columnar_fwc {\n        position: absolute;\n        z-index: 2;\n        width: 8rem;\n        height: 50px;\n        left: 50%;\n        top: 220px;\n        margin-left: -4rem;\n        background-image: linear-gradient(90deg, #C8A325 0%, #EFD05F 25%, #EFD05F 35%, #DF8343 60%, #DF8343 85%, #DD9B82 100%);\n      }\n      \n      .columnar_fwc_span {\n        position: absolute;\n        z-index: 4;\n        width: 8rem;\n        height: 50px;\n        left: 50%;\n        top: 220px;\n        margin-left: -4rem;\n        min-height: 1rem;\n      }\n      \n      .columnar_fwc_after {\n        position: absolute;\n        z-index: 2;\n        width: 8rem;\n        height: 2rem;\n        left: 50%;\n        top: 2rem;\n        margin-left: 4.5rem;\n        color: #E3B456;\n        box-sizing: border-box;\n      }\n      \n      .columnar_fwc_after_border {\n        position: absolute;\n        z-index: 2;\n        display: block;\n        width: 6px;\n        height: 100%;\n        border: 1px solid #1CBFEF;\n        border-left: none;\n        top: 0;\n        left: -0.6rem;\n      }\n      \n      .columnar_fwc_ellipse {\n        position: absolute;\n        z-index: 3;\n        width: 8rem;\n        height: 20px;\n        left: 50%;\n        top: 210px;\n        margin-left: -4rem;\n        margin-top: -10px;\n        border-radius: 50% / 50%;\n        background-image: linear-gradient(90deg, #BC933B 0%, #BA8D47 100%);\n      }\n      \n      .columnar_zwc_half {\n        position: absolute;\n        z-index: 1;\n        width: 8rem;\n        height: 20px;\n        left: 50%;\n        top: 20rem;\n        margin-left: -4rem;\n        margin-top: -10px;\n        border-radius: 50% / 50%;\n        background-image: linear-gradient(90deg, #92B385 0%, #B1D6B4 25%, #B1D6B4 35%, #32AB84 60%, #32AB84 85%, #38C5AD 100%);\n      }\n      \n      .columnar_zwc {\n        position: absolute;\n        z-index: 2;\n        width: 8rem;\n        height: 50px;\n        left: 50%;\n        top: 220px;\n        margin-left: -4rem;\n        background-image: linear-gradient(90deg, #92B385 0%, #B1D6B4 25%, #B1D6B4 35%, #32AB84 60%, #32AB84 85%, #38C5AD 100%);\n      }\n      \n      .columnar_zwc_after {\n        position: absolute;\n        z-index: 2;\n        display: block;\n        width: 8rem;\n        height: 2rem;\n        left: 50%;\n        top: 2rem;\n        margin-left: 4.5rem;\n        color: #769E8F;\n        box-sizing: border-box;\n      }\n      \n      .columnar_zwc_after_border {\n        position: absolute;\n        z-index: 2;\n        display: block;\n        width: 6px;\n        height: 95%;\n        margin-top: 5%;\n        border: 1px solid #1CBFEF;\n        border-left: none;\n        top: 0;\n        left: 0;\n      }\n      \n      .columnar_zwc_span {\n        position: absolute;\n        z-index: 4;\n        width: 8rem;\n        height: 50px;\n        left: 50%;\n        top: 220px;\n        margin-left: -4rem;\n        min-height: 1rem;\n      }\n      \n      .columnar_zwc_ellipse {\n        position: absolute;\n        z-index: 3;\n        width: 8rem;\n        height: 20px;\n        left: 50%;\n        top: 210px;\n        margin-left: -4rem;\n        margin-top: -10px;\n        border-radius: 50% / 50%;\n        background-image: linear-gradient(90deg, #2ECFB0 0%, #CBF6D6 100%);\n      }\n      \n      .columnar_quarter {\n        position: absolute;\n        z-index: 1;\n        display: block;\n        width: 15.2rem;\n        height: 2rem;\n        left: 50%;\n        top: 2rem;\n        margin-left: -19.5rem;\n        border-right: 1px solid #1CBFEF;\n        line-height: 1rem;\n      }\n      \n      .columnar_quarter:after {\n        content: '';\n        position: absolute;\n        z-index: 1;\n        right: 0;\n        margin-top: -1px;\n        display: block;\n        width: 6px;\n        height: 100%;\n        border-bottom: 1px solid #1CBFEF;\n      }\n      .columnar_quarter li {\n        width: 100%;\n      }\n      \n      .quarter_left_li,\n      .quarter_middle_li,\n      .quarter_right_li {\n        display: block;\n        width: 7.8rem;\n        float: left;\n      }\n\n      .quarter_left_li {\n        margin-top: -32px;\n        text-align: right;\n      }\n      \n      .quarter_middle_li {\n        width: 4rem;\n        margin-top: -15px;\n      }\n      \n      .quarter_middle_li li {\n        color: #1CBFEF;\n        border-top: 1px solid #1CBFEF;\n        transform: rotate(20deg);\n        -ms-transform: rotate(20deg);\n      }\n      \n      .quarter_right_li {\n        width: 3.2rem;\n        margin-top: -7px;\n      }\n      \n      .quarter_right_li li:before {\n        content: '';\n        position: absolute;\n        z-index: 1;\n        right: 0;\n        display: block;\n        margin-top: 7px;\n        width: 6px;\n        height: 10px;\n        border-top: 1px solid #1CBFEF;\n      }\n      .tb_arrowup {\n        color: #FF4873;\n      }\n      .tb_arrowdown {\n        color: #769E8F;\n      }\n      .tb_arrowup i, .tb_arrowdown i {\n        font-size: 18px;\n        width: 15px;\n      }\n      ";
    document.head.appendChild(style);
  },
  render: function render() {
    if (!this.node) {
      this.node = document.querySelector('#' + this.opt.el);

      if (this.node) {
        this.node.className = 'columnar_custom_dom';
      } else {
        return;
      }
    }

    this._render();
  },
  _render: function _render() {
    if (Number(this.opt.data.completed) < 0) {
      this.node.innerHTML = "\n        <div class=\"total_target\" style=\"top:".concat(17 - this.height, "rem;\">\n        ").concat(this.toThousands(this.opt.data.total, 20), "\u5143<br/>\n          <span style=\"color:#1CBFEF\">\u76EE\u6807\u503C</span>\n        </div>\n        <div class=\"columnar_total\" style=\"height:").concat(this.height, "rem;top:").concat(20 - this.height, "rem;\"></div>\n        <div class=\"columnar_total_span\" style=\"height:").concat(this.height - this.completedHeight, "rem;line-height:").concat(this.height - this.completedHeight, "rem;top:").concat(20 - this.height, "rem;\">\n        ").concat(Math.round((this.opt.data.incomplete / this.opt.data.total).toFixed(4) * 10000) / 100, "%\n        </div>\n        <div class=\"columnar_total_ellipse\" style=\"top:").concat(20 - this.height, "rem;\"></div>\n        <div class=\"columnar_total_after ").concat(this.ifup(), "\" style=\"height:").concat(this.height, "rem;top:").concat(20 - this.height, "rem;\">\n          \u672A\u5B8C\u6210<br/>").concat(this.toThousands(this.opt.data.incomplete, 20), "\u5143\n          <div class=\"columnar_total_after_border\"></div>\n        </div>\n        <div class=\"columnar_fwc_half\"></div>\n        <div class=\"columnar_fwc\" style=\"height:").concat(this.completedHeight, "rem;top:").concat(20 - this.completedHeight, "rem;\"></div>\n        <div class=\"columnar_fwc_span\" style=\"height:").concat(this.completedHeight, "rem;line-height:").concat(this.completedHeight, "rem;top:").concat(20 - this.completedHeight, "rem;\">\n        ").concat(Math.round((this.opt.data.completed / this.opt.data.total).toFixed(4) * 10000) / 100, "%\n        </div>\n        <div class=\"columnar_fwc_ellipse\" style=\"top:").concat(20 - this.completedHeight, "rem;\"></div>\n        <div class=\"columnar_fwc_after\" style=\"").concat(this.ifdown(), "\">\n          \u5DF2\u5B8C\u6210<br/>").concat(this.toThousands(this.opt.data.completed, 20), "\u5143\n          ").concat(this.arrowPercent(this.opt.data.completedGrowthTate), "\n          <div class=\"columnar_fwc_after_border\"></div>\n        </div>\n      ");
    } else {
      this.node.innerHTML = "\n        <div class=\"total_target\" style=\"top:".concat(17 - this.height, "rem;\">\n          ").concat(this.toThousands(this.opt.data.total, 20), "\u5143<br/>\n          <span style=\"color:#1CBFEF\">\u76EE\u6807\u503C</span>\n        </div>\n        <div class=\"columnar_total\" style=\"height:").concat(this.height, "rem;top:").concat(20 - this.height, "rem;\"></div>\n        <div class=\"columnar_total_span\" style=\"height:").concat(this.height - this.completedHeight, "rem;line-height:").concat(this.height - this.completedHeight, "rem;top:").concat(20 - this.height, "rem;").concat(this.ifshow(), "\">\n        ").concat(Math.round((this.opt.data.incomplete / this.opt.data.total).toFixed(4) * 10000) / 100, "%\n        </div>\n        <div class=\"columnar_total_ellipse\" style=\"top:").concat(20 - this.height, "rem;\"></div>\n        <div class=\"columnar_total_after ").concat(this.ifup(), "\" style=\"height:").concat(this.height - this.completedHeight, "rem;top:").concat(20 - this.height, "rem;").concat(this.ifshow(), "\">\n          \u672A\u5B8C\u6210<br/>").concat(this.toThousands(this.opt.data.incomplete, 20), "\u5143\n          <div class=\"columnar_total_after_border\"></div>\n        </div>\n        <div class=\"columnar_zwc_half\"></div>\n        <div class=\"columnar_zwc\" style=\"height:").concat(this.completedHeight, "rem;top:").concat(20 - this.completedHeight, "rem;\"></div>\n        <div class=\"columnar_zwc_span\" style=\"height:").concat(this.completedHeight, "rem;line-height:").concat(this.completedHeight, "rem;top:").concat(20 - this.completedHeight, "rem;\">\n        ").concat(Math.round((this.opt.data.completed / this.opt.data.total).toFixed(4) * 10000) / 100, "%\n        </div>\n        <div class=\"columnar_zwc_ellipse\" style=\"top:").concat(20 - this.completedHeight, "rem;\"></div>\n        <div class=\"columnar_zwc_after\" style=\"").concat(this.ifdown(), "\">\n          \u5DF2\u5B8C\u6210<br/>").concat(this.toThousands(this.opt.data.completed, 20), "\u5143\n          ").concat(this.arrowPercent(this.opt.data.completedGrowthTate), "\n          <div class=\"columnar_zwc_after_border\"></div>\n        </div>\n      ");
    }

    this.quarterRender();
  },
  quarterRender: function quarterRender() {
    if (!this.opt.data.quarterData || !(this.opt.data.quarterData instanceof Array) || this.opt.data.quarterData.length === 0) return;
    var div = document.createElement('div');
    div.className = 'columnar_quarter';
    div.style.height = this.completedHeight + 'rem';
    div.style.top = 20 - this.completedHeight + 'rem';
    var spanL = '',
      spanM = '',
      spanR = '';
    var self = this;
    var liH = 1.2;

    if (this.completedHeight > this.opt.data.quarterData.length) {
      liH = parseInt((this.completedHeight / this.opt.data.quarterData.length).toFixed(2) * 100) / 100;
    }

    this.opt.data.quarterData.reverse().map(function (o) {
      o.label = o.label || '累计营收';
      spanL += "<li style=\"height:".concat(liH, "rem;\">").concat(self.toThousands(o.value), "\u5143").concat(self.arrowPercent(o.quarterGrowthTate), "</li>");
      spanM += "<li style=\"height:".concat(liH, "rem;\">").concat(o.label, "</li>");
      spanR += "<li style=\"height:".concat(liH, "rem;\">").concat(o.name, "</li>");
    });
    div.innerHTML = "\n      <div class=\"quarter_left_li\">\n       <ul>\n       ".concat(spanL, "\n       </ul>\n      </div>\n      <div class=\"quarter_middle_li\">\n      <ul>\n      ").concat(spanM, "\n      </ul></div>\n      <div class=\"quarter_right_li\">\n      <ul>\n      ").concat(spanR, "\n      </ul></div>\n      ");
    this.node.appendChild(div);
  },
  ifshow: function ifshow() {
    if (this.opt.data.total <= this.opt.data.completed) {
      return 'display: none;';
    }

    return '';
  },
  ifup: function ifup() {
    if (this.height - this.completedHeight <= 2) {
      return 'columnar_total_min_height';
    }

    return '';
  },
  ifdown: function ifdown() {
    if (this.completedHeight > 4) {
      var ptop = (this.completedHeight - 4) / 2;
      return "height:".concat(this.completedHeight, "rem;top:").concat(20 - this.completedHeight, "rem;padding-top:").concat(ptop, "rem;");
    }

    return "height:".concat(this.completedHeight, "rem;top:").concat(20 - this.completedHeight, "rem;");
  },
  arrowPercent: function arrowPercent(n) {
    if (!n) return '';

    if (n >= 0) {
      return "<span class=\"tb_arrowup\"><i class=\"icon-arrow-up\"></i>".concat(n, "%</span>");
    } else {
      return "<span class=\"tb_arrowdown\"><i class=\"icon-arrow-down\"></i>".concat(n, "%</span>");
    }
  },
  checkData: function checkData() {
    this.opt.data.incomplete = this.opt.data.total - this.opt.data.completed;
    if (this.opt.data.incomplete < 0) this.opt.data.incomplete = 0;

    if (this.opt.data.maxTarget && this.opt.data.total < this.opt.data.maxTarget) {
      if (this.opt.data.total < this.opt.data.maxTarget / 2) {
        this.height = 8;
      } else {
        this.height = Math.round((16 * this.opt.data.total / this.opt.data.maxTarget).toFixed(2) * 100) / 100;
      }
    }

    if (this.opt.data.quarterData && this.opt.data.quarterData.length > 1) {
      this.minHeight = this.opt.data.quarterData.length;
    }

    if (this.opt.data.total <= this.opt.data.completed) {
      this.completedHeight = this.height;
    } else {
      this.completedHeight = Math.round((Math.abs(this.opt.data.completed) * this.height / this.opt.data.total).toFixed(2) * 100) / 100;

      if (this.completedHeight < this.minHeight) {
        this.completedHeight = this.minHeight;
      }

      if (this.height - this.completedHeight < 2) {
        this.completedHeight = this.height - 2;
      }
    }
  },
  toThousands: function toThousands(n, f) {
    var str = n;
    var size = f || 12;
    str = '<em style="font-size:' + size + 'px">' + Math.round(n / 1000000) / 100 + '</em>亿';
    return str;
  },
  toThousands2: function toThousands2(n) {
    var absn = Math.abs(n);
    var str = n;

    if (absn >= 1000 && absn < 10000) {
      str = Number((n / 1000).toFixed(2)) + '千';
    } else if (absn >= 10000 && absn < 1000000) {
      str = Number((n / 10000).toFixed(2)) + '万';
    } else if (absn >= 1000000 && absn < 10000000) {
      str = Number((n / 1000000).toFixed(2)) + '百万';
    } else if (absn >= 10000000 && absn < 100000000) {
      str = Number((n / 10000000).toFixed(2)) + '千万';
    } else if (absn >= 100000000 && absn < 1000000000000) {
      str = Number((n / 100000000).toFixed(2)) + '亿';
    } else if (absn >= 1000000000000) {
      str = Number((n / 1000000000000).toFixed(2)) + '万亿';
    }

    return str;
  }
};
export default columnar;