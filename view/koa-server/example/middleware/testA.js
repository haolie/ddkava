import lib from '../../lib/lib';
export default () => async (ctx, next) => {

  await next();
  var path=ctx.path;
  path= path.replace("api/v1/","")
  lib.creatJson(path,ctx.body)

};
