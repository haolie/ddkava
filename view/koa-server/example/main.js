import Koa from 'koa';
import bodyParser from 'koa-bodyparser';
import cors from 'koa-cors';

import config from './config';
import errorHandle from './middleware/errorHandle';
import testA from './middleware/testA';
import testB from './middleware/testB';
import router from './routes/index';

const serve = require('koa-static');

const app = new Koa();

app
  .use(cors())
  .use(testA())
  .use(testB())
  .use(serve('.'))
  .use(bodyParser())
  .use(errorHandle())
  .use(router.routes())

  .use(router.allowedMethods());

export default app.listen(config.port, () => {
  console.log(`App is listening on ${config.port}.`);
});
