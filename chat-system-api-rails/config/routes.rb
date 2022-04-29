Rails.application.routes.draw do
  namespace :api do
    namespace :v1 do
      resources :applications, param: :access_token, only: [:index, :show, :create, :update] do
        resources :chats, param: :number, only: [:index, :show] do
          resources :messages, param: :number, only: [:index, :update, :show] do
            collection do
              get :search
            end
          end
        end
      end
    end
  end
end
