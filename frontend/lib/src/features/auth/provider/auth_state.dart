enum AuthStatus { initial, checking,loading, authenticated, unauthenticated, error }

class AuthState {
  final AuthStatus status;
  final String? errorMessage;

  AuthState({required this.status, this.errorMessage});

  // 初始状态
  factory AuthState.initial() => AuthState(status: AuthStatus.initial);
}