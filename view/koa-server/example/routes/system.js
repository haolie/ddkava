var Mock = require('mockjs');
import { request, summary, body, tags, middlewares, path, description, okMsg, lib, errMsg } from '../../lib';

const tag = tags(['系统']);

export default class SystemRouter {
  @request('post', '/auth/system/login')
  @summary('登录')
  @tag
  @body({
    username: {type: 'string', require: true},
    password: {type: 'string', require: true}
  })
  static async login(ctx) {
    okMsg.data = {
      "token": Mock.mock(/[0-9a-zA-Z]/)
    };
    ctx.body = okMsg;
  }

  @request('get', '/auth/system/logout')
  @summary('登出')
  @tag
  static async logout(ctx) {
    okMsg.data = 1;
    ctx.body = okMsg;
  }
}
