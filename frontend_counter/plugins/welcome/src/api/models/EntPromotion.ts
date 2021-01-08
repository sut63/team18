/* tslint:disable */
/* eslint-disable */
/**
 * SUT SE Example API
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntPromotionEdges,
    EntPromotionEdgesFromJSON,
    EntPromotionEdgesFromJSONTyped,
    EntPromotionEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntPromotion
 */
export interface EntPromotion {
    /**
     * Discount holds the value of the "discount" field.
     * @type {number}
     * @memberof EntPromotion
     */
    discount?: number;
    /**
     * 
     * @type {EntPromotionEdges}
     * @memberof EntPromotion
     */
    edges?: EntPromotionEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntPromotion
     */
    id?: number;
    /**
     * PromotionName holds the value of the "promotion_name" field.
     * @type {string}
     * @memberof EntPromotion
     */
    promotionName?: string;
}

export function EntPromotionFromJSON(json: any): EntPromotion {
    return EntPromotionFromJSONTyped(json, false);
}

export function EntPromotionFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntPromotion {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'discount': !exists(json, 'discount') ? undefined : json['discount'],
        'edges': !exists(json, 'edges') ? undefined : EntPromotionEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
        'promotionName': !exists(json, 'promotion_name') ? undefined : json['promotion_name'],
    };
}

export function EntPromotionToJSON(value?: EntPromotion | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'discount': value.discount,
        'edges': EntPromotionEdgesToJSON(value.edges),
        'id': value.id,
        'promotion_name': value.promotionName,
    };
}


