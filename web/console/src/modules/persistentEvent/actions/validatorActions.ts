import { RootState, PeEdit } from '../models';
import * as ActionType from '../constants//ActionType';
import { t, Trans } from '@tencent/tea-app/lib/i18n';

type GetState = () => RootState;

export const validatorActions = {
  /** 校验当前的es地址是否正确 */
  _validateEsAddress(address: string) {
    let status = 0,
      message = '',
      hostReg = /^((http|https):\/\/)((25[0-5]|2[0-4]\d|1?\d?\d)\.){3}(25[0-5]|2[0-4]\d|1?\d?\d):([0-9]|[1-9]\d{1,3}|[1-5]\d{4}|6[0-4]\d{4}|65[0-4]\d{2}|655[0-2]\d|6553[0-5])$/;

    if (!address) {
      status = 2;
      message = t('Elasticsearch地址不能为空');
    } else if (!hostReg.test(address)) {
      status = 2;
      message = t('Elasticsearch地址格式不正确，{{scheme}}://{{addr}}:{{port}}');
    } else {
      status = 1;
      message = '';
    }
    return { status, message };
  },

  validateEsAddress() {
    return async (dispatch: Redux.Dispatch, getState: GetState) => {
      let { esAddress } = getState().peEdit;

      let result = validatorActions._validateEsAddress(esAddress);
      dispatch({
        type: ActionType.V_EsAddress,
        payload: result
      });
    };
  },

  /** 校验当前的索引名是否正确 */
  _validateIndexName(indexName: string) {
    let status = 0,
      message = '',
      reg = /^[a-z][0-9a-z_+-]*$/;
    if (!indexName) {
      status = 2;
      message = t('索引名不能为空');
    } else if (!reg.test(indexName)) {
      status = 2;
      message = t('索引名格式不正确');
    } else if (indexName.length > 60) {
      status = 2;
      message = t('索引名不能超过60个字符');
    } else {
      status = 1;
      message = '';
    }
    return { status, message };
  },

  validateIndexName() {
    return async (dispatch: Redux.Dispatch, getState: GetState) => {
      let { indexName } = getState().peEdit;
      let result = validatorActions._validateIndexName(indexName);
      dispatch({
        type: ActionType.V_IndexName,
        payload: result
      });
    };
  },

  /** 持久化事件的校验 */
  _validatePeEdit(peEdit: PeEdit) {
    let result = true;

    if (peEdit.isOpen) {
      result =
        result &&
        validatorActions._validateEsAddress(peEdit.esAddress).status === 1 &&
        validatorActions._validateIndexName(peEdit.indexName).status === 1;
    }

    return result;
  },

  validatePeEdit() {
    return async (dispatch: Redux.Dispatch, getState: GetState) => {
      let { peEdit } = getState();
      if (peEdit.isOpen) {
        dispatch(validatorActions.validateEsAddress());
        dispatch(validatorActions.validateIndexName());
      }
    };
  }
};
