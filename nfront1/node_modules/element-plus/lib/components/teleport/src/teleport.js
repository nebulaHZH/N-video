'use strict';

Object.defineProperty(exports, '__esModule', { value: true });

require('../../../utils/index.js');
var runtime = require('../../../utils/vue/props/runtime.js');

const teleportProps = runtime.buildProps({
  container: {
    type: runtime.definePropType(String),
    default: "body"
  },
  disabled: Boolean,
  style: {
    type: runtime.definePropType([String, Array, Object])
  },
  zIndex: {
    type: String,
    default: "2000"
  }
});

exports.teleportProps = teleportProps;
//# sourceMappingURL=teleport.js.map
