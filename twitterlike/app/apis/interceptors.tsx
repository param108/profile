import axios from "axios";


axios.interceptors.response.use(undefined, async (err) => {
  const { config, message } = err;
  if (!config || !config.retry) {
    return Promise.reject(err);
  }
  // retry while Network timeout or Network Error
  if (!(message.includes("timeout") || message.includes("Network Error"))) {
    return Promise.reject(err);
  }
  config.retry -= 1;
  const delayRetryRequest = new Promise((resolve) => {
    setTimeout(() => {
        console.log("retry the request", config.url);
        resolve(null);
    }, config.retryDelay || 1000);
  });
  return delayRetryRequest.then(() => axios(config));
});
