/*

    Licensed under the Apache License, Version 2.0 (the "License");
    you may not use this file except in compliance with the License.
    You may obtain a copy of the License at

        http://www.apache.org/licenses/LICENSE-2.0

    Unless required by applicable law or agreed to in writing, software
    distributed under the License is distributed on an "AS IS" BASIS,
    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
    See the License for the specific language governing permissions and
    limitations under the License.
*/

-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

DELETE FROM api_capability WHERE http_method = 'POST' and route = '/internal/api/*/steering';


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

INSERT INTO api_capability (http_method, route, capability) VALUES ('POST', '/internal/api/*/steering', 'ds-steering-write') ON CONFLICT (http_method, route, capability) DO NOTHING;
