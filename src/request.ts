import { ApiServerError } from './types/api';

export class ApiProtocolError extends Error {
    constructor(message: string, cause?: Error) {
        super(message, { cause });

        this.name = 'ApiProtocolError';
    }
}

export class ApiError extends Error {
    constructor(fromApi: ApiServerError) {
        super(fromApi.error);

        this.name = 'ApiError';
    }
}

function isApiServerError(e: unknown): e is ApiServerError {
    return typeof e === 'object' && e !== null && 'error' in e;
}

export class ApiRequest {
    private static prefix = '/api';

    static async Get<T>(path: string) {
        return await this.request<T>('GET', path);
    }

    static async Post<T>(path: string, body?: BodyInit) {
        return await this.request<T>('POST', path, body);
    }

    static async Put<T>(path: string, body?: BodyInit) {
        return await this.request<T>('PUT', path, body);
    }

    static async Patch<T>(path: string, body?: BodyInit) {
        return await this.request<T>('PATCH', path, body);
    }

    static async Delete<T>(path: string) {
        return await this.request<T>('DELETE', path);
    }

    private static async request<T>(method: string, path: string, body?: BodyInit) {
        try {
            const res = await fetch(this.prefix + path, {
                method,
                body,
                credentials: 'include',
            });

            const parsed = (await res.json()) as T;

            if (isApiServerError(parsed)) {
                throw new ApiError(parsed);
            }

            return parsed;
        } catch (e) {
            if (e instanceof ApiError) throw e;

            throw new ApiProtocolError('api request error', e as Error);
        }
    }
}
